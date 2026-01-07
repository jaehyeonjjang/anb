import { ref, onMounted, onUnmounted } from "vue"

export default function() {
    const w = ref(0)
    const h = ref(0)

    const setSize = () => {         
        w.value = window.innerWidth
        h.value = window.innerHeight
    }

    const height = (adjust: number) => {
        return `${h.value - adjust}px`
    }

    const width = (adjust: number) => {
        return `${w.value - adjust}px`
    }
    
    onMounted(() => {
        setSize()

        window.addEventListener('resize', setSize)
    })

    onUnmounted(() => {
        window.removeEventListener('resize', setSize)
    })

    return {
        width,
        height
    }
}    
