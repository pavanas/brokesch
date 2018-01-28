<template>
<ul class="vue-countdown">
  <input type="hidden" name="" :value="deadline" />
    <li>
        <p class="digit">{{ days | twoDigits }}</p>
        <p class="text">days</p>
    </li>

    <li>
        <p class="digit">{{ hours | twoDigits }}</p>
        <p class="text">hours</p>
    </li>

    <li>
        <p class="digit">{{ minutes | twoDigits }}</p>
        <p class="text">Min</p>
    </li>

    <li>
        <p class="digit">{{ seconds | twoDigits }}</p>
        <p class="text">Sec</p>
    </li>
</ul>
</template>

<script>
import Vue from 'vue'

Vue.filter('twoDigits', (value) => {
    if ( value.toString().length <= 1 ) {
        return '0'+value.toString()
    }
    return value.toString()
})

let interval = null;

export default {
    props: ['deadline'],

    data() {
        let d_end = new Date();
        d_end.setHours(d_end.getHours() + 10);
        return {
            now: Math.trunc((d_end).getTime() / 1000),
            date: null
        }
    },

    mounted() {
        
        this.date = Math.trunc(Date.parse(this.deadline) / 1000)

        interval = setInterval(() => {
            let tt = new Date();
            tt.setHours(tt.getHours() + 10);
            // this.now = Math.trunc((new Date()).getTime() / 1000)
            this.now = Math.trunc((tt).getTime() / 1000)
        }, 1000)
    },

    computed: {
        seconds() {
            return Math.trunc(this.date - this.now) % 60
        },

        minutes() {
            return Math.trunc((this.date - this.now) / 60) % 60
        },

        hours() {
            return Math.trunc((this.date - this.now) / 60 / 60) % 24
        },

        days() {
            return Math.trunc((this.date - this.now) / 60 / 60 / 24)
        }
    },

    watch: {
      now: function(value){
          this.diff = this.date - this.now;
          if(this.diff <= 0 || this.stop){
              this.diff = 0;
              // Remove interval
              clearInterval(interval);
              this.$emit('countend');
          }
      }

}
}
</script>
