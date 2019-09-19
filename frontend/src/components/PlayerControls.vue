<template>
    <div class="playerControls">
        <div class="controlsContainer">
            <div class="controls">
                <div class="controlRestart">
                    <button class="controlButtonRestart"></button>
                </div>
                <div class="controlPlay">
                    <button class="controlButtonPlay" v-on:click.prevent="play"></button>
                </div>
                <div class="controlStop">
                    <button class="controlButtonStop" v-on:click.prevent="stop"></button>
                </div>
            </div>
        </div>
        
    </div>
</template>

<script>
export default {
    name: "PlayerControls",
    props: ['player'],
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
.controls {
    padding: 15px 0px;
    display: flex;
    flex-direction: row;
    justify-content: center;
}

.controlButtonRestart {
    width: 30px;
    height: 30px;
    border-style: solid;
    border-width: 15px 15px 15px 15px;
    border-color: rgb(200, 200, 200);
    background: transparent;
    box-sizing: border-box;
}

.controlButtonPlay {
    width: 30px;
    height: 30px;
    border-style: solid;
    border-width: 15px 15px 15px 15px;
    border-color: rgb(200, 200, 200);
    background: transparent;
    box-sizing: border-box;
}

.controlButtonStop {
    width: 30px;
    height: 30px;
    border-style: solid;
    border-width: 15px 15px 15px 15px;
    border-color: rgb(200, 200, 200);
    background: transparent;
    box-sizing: border-box;
}

.controlRestart,
.controlPlay,
.controlStop {
    padding: 0px 20px;
}
</style>