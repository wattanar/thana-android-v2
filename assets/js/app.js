jQuery(document).ready(function($) {
	home();
});

function home() {
	$('a#app-home').hide();
	setTimeout(function(){
		$('a#app-home').fadeIn('fast');
	},1000);
	$('.tile-wrap').html("");
	$('.movie_list').html("");
	$('.progress').show();

	get_theater()
		.done(function(data) {
			$.each(data, function(index, val) {
				$('.tile-wrap').append(
					'<div class="tile">' +
						'<div class="tile-side pull-left"> <span class="icon icon-2x">theaters</span> </div>' +
			            '<div class="tile-inner" onclick="return listmovies('+val.Link+', \''+val.Title+'\')">'+ val.Title +' </div>' +
			        '</div>'
			    );
			});
			$('.progress').hide();
		})
		.fail(function() {
			$('.tile-wrap').html("error : disconnected from network!");
		    $('.progress').hide();
		});
}

function listmovies(id, T_title) {

	$('.tile-wrap').html("");
	$('.movie_list').html("");
	$('.progress').show();

	$('.movie_list').append(
		'<div class="tile-wrap">' +
			'<div class="tile">' +
				'<div class="tile-side pull-left"> <span class="icon icon-2x">theaters</span> </div>' +
		        '<div class="tile-inner">'+ T_title +' </div>' +
		    '</div>' +
		'</div>'
	);

	get_movie(id)
		.done(function(data) {
			$.each(data, function(index, val) {
				
				var str_time = '';
				var temp_time = val.Time.match(/.{1,5}/g);

				$.each(temp_time, function(index, val) {
					 str_time += ('<span class="label label-brand">' + val +'</span> ');
				});

				$('.movie_list').append(
					'<div class="card">' +
						'<aside class="card-side card-side-img pull-left" style="width:150px">' +
							'<img width="100%" height="250px" src="'+val.Thumb+'"> ' +
						'</aside> ' +
						'<div class="card-main"> ' +
							'<div class="card-inner"> ' +
								'<p class="card-heading">'+val.Title_EN+'</p> ' +
								'<p class="margin-bottom-lg">'+val.Title_TH+'</p> ' +
							'</div> ' +
							'<div class="card-action" style="padding:20px;"> เวลาฉาย : '+str_time+'</div>' +
						'</div> ' +
					'</div> ' 
			    );
			});
			$('.progress').hide();
		})
		.fail(function() {
			$('.tile-wrap').html("error : disconnected from network!");
		    $('.progress').hide();
		});
}

function get_movie(id) {
	return $.ajax({
		url : '/t/'+id,
		type : 'get',
		dataType : 'json'
	});
}

function get_theater() {
	return $.ajax({
		url : '/theater',
		type : 'get',
		dataType : 'json'
	});
}	

