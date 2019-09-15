<template>
    <div class="layoutMiddle">
        <MiddleLeft />
        <MiddleRight />
    </div>
</template>

<script>
import MiddleLeft from './Middle/Left'
import MiddleRight from './Middle/Right'

export default {
    name: "Middle",
    components: {
        MiddleLeft,
        MiddleRight
    }
}
</script>