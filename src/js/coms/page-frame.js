/**
 * @desc 页面布局通用组件, 包含页头、页脚和内容
 * @author {{ author }}
 * @email {{ email }}
 * @create {{ create }}
 */
import headerComponents from "./header"
import footerComponents from "./footer"
export default {
    data() {
        return {
            page: {
                width: innerWidth,
                height: innerHeight
            },
        }
    },
    template: `<div id="page-content">
        <!-- -------------------- 页头 -------------------- -->
        <page-header></page-header>
        
        <!-- -------------------- 内容 -------------------- -->
        <slot></slot>
        
        <!-- -------------------- 页脚 -------------------- -->
        <page-footer></page-footer>
    </div>`,
    components: {
        'page-header': headerComponents,
        'page-footer': footerComponents,
    },
    mounted() {
    }
}
