
<div class="filter">
    <!--span class="filtercaption">Filter</span><br>-->
    <table>
        <tr>
            <td class="caption">Search</td>
            <td class="input"><input class="search"></td>
        </tr>
    </table>
    
</div>

<div id="bookmarksContainer"></div>

<script type="text/javascript">
    t = 1;
    bookmarks = JSON.parse('<?php echo json_encode($this->par['bookmarks'],TRUE);?>');

    var domainSet = new Set();
    <?php foreach($this->par['bookmarks'] as $bookmark){
        ?>domainSet.add('<?php echo $bookmark['domain'];?>'); <?php
    }
    ?>
    
    initializeFiler();
    printBookmarkList();

    $( document ).ready(function() {
        $(".filter .search").on('input', function() {
            
            if(t){
                clearTimeout(t);
            }
            searchString = $(this).val();
            t = setTimeout(
                (function(){
                    changeFiler(searchString);
                    printBookmarkList();
                }), 200);
            
        });
    });
    
    
</script>			