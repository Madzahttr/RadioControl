<template>
    <div class="songInfo">
        <div class="controlsContainer">
            <div class="songName">
                {{ currentSong.name }}
            </div>
            <div class="controlsContainerRow">
                <div class="songArtist">
                    {{ currentSong.artist }}
                </div>
                <div class="bull">&bull;</div>
                <div class="songAlbum">
                    {{ currentSong.album }}
                </div>
                <div class="bull">&bull;</div>
                <div class="songYear">
                    {{ currentSong.releaseYear }}
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "SongInfo",
    props: ["currentSong"]
}
</script>

<style scoped>
.controlsContainer {
    display: flex;
    flex-direction: column;
}

.controlsContainerRow {
    display: flex;
    flex-direction: row;
    padding: 0 20%;
    padding-top: 6px;
}

.songName {
    padding: 0 20%;
    font-size: 23px;
}

.songArtist,
.songAlbum,
.songYear {
    font-size: 18px;
    padding: 0px 5px;
    color: rgb(200, 200, 200)
}

.songArtist {
    padding-left: 0px;
}

.bull {
    height: 30px;
    font-size: 18px
}
</style>