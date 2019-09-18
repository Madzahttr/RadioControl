<template>
    <div class="layoutControls">
        <PlayerControls />
        <SongInfo v-bind:currentSong="currentSong" />
        <PlayerOptions />
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
    props: ['currentSong']
}
</script>