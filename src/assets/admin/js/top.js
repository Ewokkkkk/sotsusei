$(window).on('load',function(){
    $("#splash-logo").delay(1200).fadeOut('slow');
  
    $("#splash").delay(1500).fadeOut('slow',function(){
    
       $('body').addClass('appear');
        
    });
   
   $('.splashbg').on('animationend', function() {    
    });
});

$(function() {
    $(window).scroll(function() {
      $(".scroll").each(function() {
        var scroll = $(window).scrollTop();
        var blockPosition = $(this).offset().top;
        var windowHeihgt = $(window).height();
        if (scroll > blockPosition - windowHeihgt + 300) {
          $(this).addClass("blockIn");
        }
      });
    });
  });