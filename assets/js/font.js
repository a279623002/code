!function(n){
    var  e=n.document,
        t=e.documentElement,
        i=720,
        d=i/100,
        o="orientationchange"in n?"orientationchange":"resize",
        a=function(){
            var n=t.clientWidth||320;n>720&&(n=720);
            t.style.fontSize=n/d+"px"
        };
    e.addEventListener && (n.addEventListener(o, a, !1), e.addEventListener("DOMContentLoaded", a, !1))
    
    
    
}(window);

$(function () { 
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