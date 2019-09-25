<template>
    <div class="playerOptions">
        <div class="controlsContainer">
            <div class="buttons">
                <div class="playTime">
                    <div class="playTimeSpan">{{ this.currentTime }} / {{ this.durationTime }}</div>
                </div>
                <div class="settings">
                    <button class="buttonSettings"></button>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "PlayerOptions",
    props: ['currentTime', 'durationTime']
}
</script>

<style scoped>
.buttons {
    padding: 15px 0px;
    display: flex;
    flex-direction: row;
    justify-content: center;
}

.playTime {
    height: 30px;
}

.playTimeSpan {
    height: 30px;
    font-size: 18px;
    padding-top: 5px;
    padding-bottom: 7px;
    padding-left: 5px;
    padding-right: 20px;
    color: rgb(200, 200, 200);
    text-align: center;
}

.buttonSettings {
    width: 30px;
    height: 30px;
    border-style: solid;
    border-width: 15px 15px 15px 15px;
    border-color: rgb(200, 200, 200);
    background: transparent;
    box-sizing: border-box;
}
</style>