<template>
    <div id="app">
        <Top />
        <Middle />
        <Bottom />
        <Controls />
    </div>
</template>

<script>
import './assets/css/main.css'
import Top from './components/layouts/Top'
import Middle from './components/layouts/Middle'
import Bottom from './components/layouts/Bottom'
import Controls from './components/layouts/Controls'

export default {
    name: "app",
    components: {
        Top,
        Middle,
        Bottom,
        Controls
    }
};
</script>