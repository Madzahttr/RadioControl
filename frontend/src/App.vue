<template>
    <div id="app">
        <Top />
        <Middle />
        <Bottom />
        <Controls v-bind:currentSong="currentSong" v-bind:player="player" v-bind:currentTime="currentTime" v-bind:durationTime="durationTime" v-bind:percentComplete="percentComplete" @stop="stop" @play="play" />
		<audio ref="audiofile" :src="currentSong.file" preload="auto"></audio>
    </div>
</template>

<script>
import './assets/css/main.css'
import Top from './components/layouts/Top'
import Middle from './components/layouts/Middle'
import Bottom from './components/layouts/Bottom'
import Controls from './components/layouts/Controls'

const convertTimeHHMMSS = (val) => {
	let hhmmss = new Date(val * 1000).toISOString().substr(11, 8);

	return hhmmss.indexOf("00:") === 0 ? hhmmss.substr(3) : hhmmss;
};

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
                name: "Faded",
                artist: "Alan Walker",
                album: "Faded",
                releaseYear: "2015",
                genre: "EDM",
                file: "https://dl.madzahttr.com/Faded.mp3"
            },
            player: {
                audio: undefined,
                currentSeconds: 0,
                durationSeconds: 0,
                loaded: false,
                playing: false
			},
			songs: {

			}
        }
    },
    computed: {
		currentTime() {
			return convertTimeHHMMSS(this.player.currentSeconds);
		},
		durationTime() {
			return convertTimeHHMMSS(this.player.durationSeconds);
		},
		percentComplete() {
			return parseInt(this.player.currentSeconds / this.player.durationSeconds * 100);
		}
    },
    methods: {
		load() {
			if (this.player.audio.readyState >= 2) {
				this.player.loaded = true;
				this.player.durationSeconds = parseInt(this.player.audio.duration);
				return this.player.playing = true;
			}

			throw new Error('Failed to load sound file.');
        },
        play() {
			this.player.playing = true;
			this.player.audio.play();
        },
		stop() {
			this.player.playing = false;
			this.player.audio.pause();
			this.player.audio.currentTime = 0;

		},
		update() {
			this.player.currentSeconds = parseInt(this.player.audio.currentTime);
		}
	},
	mounted() {
		this.player.audio = this.$el.querySelectorAll('audio')[0];
		this.player.audio.addEventListener('timeupdate', this.update);
		this.player.audio.addEventListener('loadeddata', this.load);
		this.player.audio.addEventListener('pause', () => { this.player.playing = false; });
		this.player.audio.addEventListener('play', () => { this.player.playing = true; });
		this.player.audio.addEventListener('ended', () => {
			this.currentSong.file = "https://dl.madzahttr.com/Unity.mp3"
		});
	}
};
</script>