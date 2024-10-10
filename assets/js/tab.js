$(function () {
    $("img.lazyload").lazyload();
    $('.tab a').on('click', function () {
        $(this).addClass('active').siblings().removeClass('active');
        var index = $(this).index();
        $('.content .box').eq(index).addClass('selected').siblings().removeClass('selected');
        $('html, body').animate({
            scrollTop: 0 // 0表示滚动条垂直位置的起始点
        }, 'slow'); // 'slow'定义动画执行的速度
    })
    $('.btn').on('click', function () {
        $('.tab').toggle(function () { }, function () { });
        if ($(this).hasClass('close')) {
            $(this).removeClass('close').html('+')
            $('.content').animate({
                left: '0',
                width: '100%'
            }, 'slow'); // 'slow'定义动画执行的速度
        } else {
            $(this).addClass('close').html('—')
            $('.content').animate({
                left: '30%',
                width: '70%'
            }, 'fast'); // 'slow'定义动画执行的速度
        }
    })

    $(".back").attr("href", "javascript:back()");
})

function back() {
    // 获取当前页面的完整URL  
    var currentUrl = window.location.href;

    currentUrl = currentUrl.split('/').slice(0, -1).join('/');

    var newPath = currentUrl
    // 如果新的路径为空字符串（即当前URL是根路径），则新路径应设为'/'  
    if (newPath === '') {
        newPath= '/';
    }

    window.location.href = newPath;

}

