<template>
    <div class="playerOptions">
        <div class="controlsContainer">
            <div class="buttons">
                <div class="playTime">
                    <span class="playTimeSpan">0:00 / 4:20</span>
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
    name: "PlayerOptions"
}
</script>

<style scoped>
.buttons {
    padding: 15px 0px;
    display: flex;
    flex-direction: row;
    justify-content: center;
}

.play

.playTimeSpan {
    height: 30px;
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