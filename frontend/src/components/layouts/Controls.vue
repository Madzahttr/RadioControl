<template>
    <div>
        <div class="playProgress">
            <div class="progressBar" :style="{ width: this.percentComplete + '%' }"></div>
        </div>
        <div class="layoutControls">
            <PlayerControls v-bind:player="player" @stop="stop" @play="play" />
            <SongInfo v-bind:currentSong="currentSong" />
            <PlayerOptions v-bind:currentTime="currentTime" v-bind:durationTime="durationTime" />
        </div>
    </div>
</template>

<script>
import PlayerControls from '../PlayerControls'
import SongInfo from '../SongInfo'
import PlayerOptions from '../PlayerOptions'

export default {
    name: "Controls",
    components: {
        PlayerControls,
        SongInfo,
        PlayerOptions
    },
    props: ['currentSong', 'player', 'currentTime', 'durationTime', 'percentComplete'],
    methods: {
        stop() {
            this.$emit('stop')
        },
        play() {
            this.$emit('play')
        }
    }
}
</script>

<style scoped>
.playProgress {
    height: 5px;
    width: 100vw;
    background: rgb(40, 40, 40)
}

.progressBar {
    background: #673699;
    height: 5px;
    width: 0.25%;
}
</style>