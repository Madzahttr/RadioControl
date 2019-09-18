<template>
    <div id="app">
        <Top />
        <Middle />
        <Bottom />
        <Controls v-bind:currentSong="currentSong" />
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
    },
    data() {
        return {
            currentSong: {
                name: "Africa",
                artist: "Toto",
                album: "Toto IV",
                releaseYear: "1982",
                genre: "Classic Rock"
            }
        }
    }
};
</script>