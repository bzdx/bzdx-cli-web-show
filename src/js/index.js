/**
 * @desc 项目入口文件, 网站首页内容脚本文件
 * @author {{ author }}
 * @email {{ email }}
 * @create {{ create }}
 */
import "../css/index.styl"
new Vue({
    el: "#page",
    data: {
        title: "hello bzdx"
    },
    template: `<div>
        {{ title }}
    </div>`
});
