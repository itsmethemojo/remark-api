function Remark(config) {
    this.readConfig(config);
    this.initialize();
    this.listen();
}

Remark.prototype.readConfig = function (config) {
    this.apiUrl = config.apiUrl;
    this.containerDivId = config.containerDivId ? "#" + config.containerDivId : '#bookmarks';
    this.filterInputId = config.filterInputId ? "#" + config.filterInputId : '#filter';
    this.remarkSelectId = config.remarkSelectId ? "#" + config.remarkSelectId : '#remark';
    this.clickSelectId = config.clickSelectId ? "#" + config.clickSelectId : '#click';
    this.firstEntriesCount = 30;
    this.wto = 0;
    this.filter = $(this.filterInputId).val();
    this.remarks = $(this.remarkSelectId).val();
    this.clicks = $(this.clickSelectId).val();
    this.bookmarks = localStorage.getObject("bookmarks") || new Array();
    this.maxCount = this.getUrlParameter("items");
}

Remark.prototype.listen = function () {
    var self = this;

    $(self.filterInputId).on('input', function (event) {
        self.filter = $(this).val().toLowerCase().trim();

        clearTimeout(self.wto);
        self.wto = setTimeout(function () {
            self.printBookmarks();
        }, 500);
    });

    $(self.remarkSelectId).on('change', function (event) {
        self.remarks = $(this).val();

        clearTimeout(self.wto);
        self.wto = setTimeout(function () {
            self.printBookmarks();
        }, 0);
    });

    $(self.clickSelectId).on('change', function (event) {
        self.clicks = $(this).val();

        clearTimeout(self.wto);
        self.wto = setTimeout(function () {
            self.printBookmarks();
        }, 0);
    });
    //TODO extract ids
    $("#toggle-insert").click(function () {
        $("#insert").toggleClass("toggled");
        if ($(this).val() === "+") {
            $(this).val("-");
        } else {
            $(this).val("+");
        }
    });

    $("#insert").on("keydown", function (event) {
        if (event.which == 13) {
            // TODO whats this
            var postUrl = self.apiUrl + "remark/?url=" + $(this).val();
            var $thisObject = $(this);
            $.getJSON(postUrl, function () {
                $thisObject.val("");
                $thisObject.toggleClass("toggled");
                self.refresh();
            });
        }

    });

}

Remark.prototype.initialize = function () {
    if (this.bookmarks.length !== 0) {
        //just print the old stuff at first
        this.printBookmarks();
    }
    this.refresh();
}

Remark.prototype.refresh = function () {
    var self = this;
    var jsonUrl = self.apiUrl;
    $.getJSON(jsonUrl, function (bookmarks) {
        self.bookmarks = bookmarks;
        localStorage.setObject("bookmarks", bookmarks);
        self.printBookmarks();
    }).fail(function (jqXHR) {
        if (jqXHR.status === 401) {
            self.login();
        }
    });
}

Remark.prototype.printBookmarks = function () {
    var self = this;
    var html = "";
    var bookmarksHtmlCreated = 0;
    previousId = 0;
    for (var i = 0; i < self.bookmarks.length; i++) {
        if (this.isBookmarkFiltered(self.bookmarks[i], i === 0 ? {"id": null} : self.bookmarks[previousId])) {
            continue;
        }
        if(self.maxCount !== null && self.maxCount == bookmarksHtmlCreated){
            break;
        }
        bookmarksHtmlCreated++;
        if (bookmarksHtmlCreated === self.firstEntriesCount) {
            $(self.containerDivId).html('<ul class="items">' + html + '</ul>');
        }
        html += this.printBookmark(self.bookmarks[i]);
        previousId = i;
    }
    $(self.containerDivId).html('<ul class="items">' + html + '</ul>');
    $("span.title a").click(function () {
        $anker = $(this);
        $.getJSON(
                self.apiUrl + "click/" + $anker.closest("li").data("id") + "/",
                function (result) {
                    self.refresh();
                }
        );
    });

}

Remark.prototype.printBookmark = function (bookmark) {
    return '<li data-id="' + bookmark['id'] + '" data-remark="' + bookmark['remarks'] + '" data-click="' + bookmark['clicks'] + '">' +
            '<span class="date">' + this.extractDate(bookmark['created']) + '</span>' +
            '<span class="time">' + this.extractTime(bookmark['created']) + '</span>' +
            '<div class="icon remark level' + this.getRemarkVisibility(bookmark['remarks']) + '"><div></div><div></div><div></div><div></div></div>' +
            '<div class="icon click level' + this.getClickVisibility(bookmark['clicks']) + '"><div></div><div></div><div></div><div></div></div>' +
            '<span class="title">' +
            '<a target="_blank" href="' + bookmark['url'] + '">' +
            (bookmark['customtitle'] === "" ? bookmark['title'] : bookmark['customtitle']) +
            '</a></span>' +
            '<br><span class="domain">' + bookmark['domain'] + '</span>' +
            '</li>';
}

Remark.prototype.isBookmarkFiltered = function (bookmark, lastBookmark) {
    if (lastBookmark['id'] === bookmark['id']) {
        return true;
    }

    if (this.remarks !== "" && this.remarks !== "=0" && this.remarks > bookmark['remarks']) {
        return true;
    }

    if (this.remarks === "=0" && bookmark['remarks'] > 0) {
        return true;
    }

    if (this.clicks !== "" && this.clicks !== "=0" && this.clicks > bookmark['clicks']) {
        return true;
    }

    if (this.clicks === "=0" && bookmark['clicks'] > 0) {
        return true;
    }


    if (this.filter === "") {
        return false;
    }

    //determine if single or multi term
    if (-1 === this.filter.indexOf(" ")) {
        if (
                -1 !== bookmark['title'].toLowerCase().indexOf(this.filter)
                || -1 !== bookmark['customtitle'].toLowerCase().indexOf(this.filter)
                || -1 !== bookmark['url'].toLowerCase().indexOf(this.filter)
                ) {
            return false;
        }
    } else {
        var searchTerms = this.filter.split(" ");
        for (var i = 0; i < searchTerms.length; i++) {
            if (searchTerms[i] === "") {
                continue;
            }
            if (
                    -1 === bookmark['title'].toLowerCase().indexOf(searchTerms[i])
                    && -1 === bookmark['customtitle'].toLowerCase().indexOf(searchTerms[i])
                    && -1 === bookmark['url'].toLowerCase().indexOf(searchTerms[i])
                    ) {
                return true;
            }
        }
        return false
    }


    return true;
}

Remark.prototype.extractDate = function (unixTimestamp) {
    var a = new Date(unixTimestamp * 1000);
    var year = a.getFullYear();
    var month = a.getMonth() < 9 ? "0" + (a.getMonth() + 1) : (a.getMonth() + 1);
    var date = a.getDate() < 10 ? "0" + a.getDate() : a.getDate();
    return date + "." + month + "." + year;
}

Remark.prototype.extractTime = function (unixTimestamp) {
    var a = new Date(unixTimestamp * 1000);
    var hour = a.getHours() < 10 ? "0" + a.getHours() : a.getHours();
    var minute = a.getMinutes() < 10 ? "0" + a.getMinutes() : a.getMinutes();
    return hour + ":" + minute;
}

Remark.prototype.getRemarkVisibility = function (count) {
    switch (parseInt(count)) {
        case 0:
            return 0;
        case 1:
            return 2;
        case 2:
            return 4;
        case 3:
            return 6;
    }
    return 8;
}

Remark.prototype.getClickVisibility = function (count) {
    switch (parseInt(count)) {
        case 0:
            return 0;
        case 1:
            return 1;
        case 2:
            return 2;
        case 3:
            return 3;
    }
    if (count <= 6)
        return 4;
    if (count <= 10)
        return 5;
    if (count <= 15)
        return 6;
    if (count <= 20)
        return 7;

    return 8;
}

Remark.prototype.login = function () {
    //no authorization -> no bookmarks cache
    console.log("not logged in");
    localStorage.setObject("bookmarks", null);
    window.location.href = "login.php";
}

Remark.prototype.getUrlParameter = function (key) {
  var regexS = "[\\?&]"+key+"=([^&#]*)";
  var regex = new RegExp( regexS );
  var results = regex.exec( location.href );
  return results == null ? null : results[1];
}

Storage.prototype.setObject = function (key, value) {
    this.setItem(key, JSON.stringify(value));
}

Storage.prototype.getObject = function (key) {
    return JSON.parse(this.getItem(key));
}