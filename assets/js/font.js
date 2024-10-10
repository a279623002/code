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

    var urlArr = currentUrl.split('/')
    if (urlArr[urlArr.length - 1] === 'index.html' || urlArr[urlArr.length - 1] === '') {
        urlArr.pop();
        urlArr.pop();
    } else {
        urlArr.pop();
    }
    var newPath = urlArr.join('/');

    window.location.href = newPath;

}