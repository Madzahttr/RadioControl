<template>
    <div class="layoutBottom">
        <BottomLeft />
        <BottomRight />
    </div>
</template>

<script>
import BottomLeft from './Bottom/Left'
import BottomRight from './Bottom/Right'

export default {
    name: "Bottom",
    components: {
        BottomLeft,
        BottomRight
    }
}
</script>