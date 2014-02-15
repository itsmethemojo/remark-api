<!DOCTYPE html>

<html lang="en">
    <head>
        <meta charset="utf-8" />
        <title><?php echo $title; ?></title>
			<link rel="stylesheet" type="text/css" href="style/main.css">
			<link href="img/favicon.ico" rel="shortcut icon" type="image/x-icon">
    </head>
    <body>
	
		<div>
			<table>

			<?php 
            if ($items):
            foreach ($items as $item): ?>

				<?php 
            if ($item['showdate']):?>
			
					<tr><td colspan="5"><br> <?php echo $item['date']; ?><br><br></td></tr>
			
				<?php endif; ?>

				<tr>
					<td class="remark-time">
						<span><?php echo $item['time']; ?><span>
					</td>
		
					<td>
						<img class="icon-padding click-icon-<?php echo $item['remarkclass']; ?>"

						<?php 
		            if ($item['bookmarkcount']==0):?>
							src="img/remark-0.png" 
						<?php else: ?>
							src="img/remark-1.png" 
						<?php endif; ?>
				
						>
						<br>
					</td>
			
			
			
					<td>
						<img 
						class="icon-padding click-icon-<?php echo $item['clickclass']; ?>" 
				
						<?php 
		            if ($item['clickcount']==0):?>
							src="img/click-0.png" 
						<?php else: ?>
							src="img/click-1.png" 
						<?php endif; ?>
				
						>
						<br>
					</td>
			
					<td>
						<a href="#');"><img class="icon-padding icon-no-count" src="img/edit.png"></a>
					</td>
					<td>
						<a class="remark-name" href="index.php?openid=<?php echo $item['id']; ?>" target="_blank"><?php echo $item['title']; ?></a>
						<br>
						<span class="domain"><?php echo $item['domain']; ?></span>
				</td>
				</tr>


		<?php 
            endforeach;
            else: ?>

        <h1>You have 0 Bookmarks.</h1>

        <?php endif; ?>

			</table>
		</div>








		
    </body>
</html>
