import type {
    Directive, App } from 'vue';
const debounce = function (func: any, delay: any) {

    let timer: any = null
    return function () {

        if (timer) clearTimeout(timer)
        timer = null
        let self = this
        let args = arguments
        timer = setTimeout(() => {

            func.apply(self, args)
        }, delay)
    }
}
const infinite: Directive = {

    mounted (el: any, binding: any, vnode: any) {
        //const { expand } = binding.modifiers
        const expand = true
        // Use richer functions , The instructions supporting the parent component act on the specified child component
        if (expand) {

            /** * target The goal is DOM The class name of the node * distance Reduce the distance threshold that triggers loading , Unit is px * func How to trigger * delay Anti shake delay , Unit is ms * load-more-disabled Whether to
             * disable infinite loading */
            let func = binding.value
            let distance = 0
            let target = '.el-scrollbar__wrap'
            let delay = 200            
            if (typeof target !== 'string') return
            let targetEl = el.querySelector(target)
            if (!targetEl) {
                console.log(' Container not found ')
                return
            }
            binding.handler = function () {
                const {
                    scrollTop, scrollHeight, clientHeight } = targetEl
                let disabled = el.getAttribute('infinite-disabled')
                disabled = vnode[disabled] || disabled

                if (scrollHeight <= Math.ceil(scrollTop + clientHeight + distance)) {

                    if (disabled == 'true') return
                    func && func()
                }
            }
            targetEl.addEventListener('scroll', binding.handler)
        } else {

            binding.handler = debounce(function () {
                const {
                    scrollTop, scrollHeight, clientHeight } = el
                if (scrollHeight === scrollTop + clientHeight) {

                    binding.value && binding.value()
                }
            }, 200)
            el.addEventListener('scroll', binding.handler)
        }
    },
    unmounted (el: any, binding: any) {

        let {
            arg } = binding
        // Use richer functions , The instructions supporting the parent component act on the specified child component 
        if (arg === 'expand') {

            /** * target The goal is DOM The class name of the node * offset Distance threshold that triggers loading , Unit is px * method How to trigger * delay Anti shake delay , Unit is ms */
            const {
                target } = binding.value
            if (typeof target !== 'string') return
            let targetEl = el.querySelector(target)
            targetEl && targetEl.removeEventListener('scroll', binding.handler)
            targetEl = null
        } else {

            el.removeEventListener('scroll', binding.handler)
            el = null
        }
    }
};

export function setupInfiniteDirective(app: App) {
    app.directive('infinite', infinite);
}

export default infinite;
