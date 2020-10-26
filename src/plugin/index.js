let plugin = {
    install(Vue, options = {}) {
        Vue.directive("limit", (el, data) => {
            let str = el.innerText
            //修饰符 upper
            if (data.modifiers.upper) {
                el.innerText = str.toUpperCase()
            }
            if (str.length > data.value) {
                el.innerText = str.substring(0, data.value) + "..."
            }
        })
        //创建中央 $bus
        const Bus = new Vue()
        Vue.prototype.$bus = Bus
    }
}
export default plugin