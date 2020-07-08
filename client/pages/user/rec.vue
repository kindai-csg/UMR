<template>
  <v-layout column>
    <v-card v-for="n in number" :key=number>
      <v-card-title>Camera {{ n }}</v-card-title>
      <v-card-text>
        <img v-bind:id="'player'+n" src=""/>
      </v-card-text>
    </v-card>
  </v-layout>
</template>

<script>
import io from 'socket.io-client';

export default {
  layout: 'user',
  data() {
    return {
      number: 2,
      fps: 30,
    }
  },
  mounted() {
    this.socket = io({
        path: '/rec/socket.io/',
      })

    const interval = 1000 / this.fps
    for (let i = 1; i <= this.number; i++) {
      setInterval(() => {
        this.socket.emit('live_'+String(i))
      }, 70)

      this.socket.on('live_'+String(i), message => {
        const player = document.getElementById("player"+String(i))
        player.setAttribute("src", "data:image/jpg;base64,"+message.data)
      })
    }
  }
}
</script>