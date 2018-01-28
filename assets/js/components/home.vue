<template>
  <div>
    <section class="hero is-info is-medium">
      <div class="hero-body">
        <div class="container has-text-centered">
          <h1 class="title">
            Next 5
          </h1>
          <h2 class="subtitle is-2">
            <strong>Visualizing the last 5 races for easier betting!</strong>
          </h2>
        </div>
      </div>
    </section>
    <section class="section is-medium">
      <div class="container has-text-centered">
        <div class="titled">
          <strong class="title is-1">Races</strong>
        </div>
      </div>
      <div class="container">
        <div class="columns is-mobile is-multiline">
          <div class="column is-one-fifth" v-for="meeting in meetings">
            <div class="card">
              <div class="card-image has-text-centered">
                <figure v-on:click="showMeetingRaces" class="image" :class="['is-'+meeting.type]" :id="[meeting.id]">
                  <img :src="['http://localhost:3000/assets/images/'+ meeting.type +'.png']" />
                </figure>
              </div>
              <div class="card-content has-text-centered">
                <strong>{{ meeting.type }}</strong>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="container races-content">
        <div v-for="race in races" class="level" :class="['is-'+race.type]">
          <div class="level-item"></div>
          <div class="level-item">
            <div>
              <p class="title is-1"><strong>Race #{{ race.number }}</strong></p>
            </div>
          </div>
          <div class="level-item" style="min-width: 30%; max-width: 30%;">
            <div>
              <p class="is-spaced heading">Time remaining</p>
              <p class="title"><countdown :deadline="race.ends_at" v-on:countend="updateRaces()"></countdown></p>
            </div>
          </div>
          <div class="level-item">
            <div>
              <p><router-link :to='{name: "showRace", params: {id: race.id}}' class="button is-medium">
                <span class="icon"><i class="fa fa-eye"></i></span>
                <span>View</span>
              </router-link></p>
              <p><a class="button is-bet is-medium">
                <span class="icon"><i class="fa fa-money"></i></span>
                <span>Bet</span>
              </a></p>
            </div>
          </div>
        </div>
      </div>
    </section>
    <footer class="footer">
      <div class="container">
        <div class="content has-text-centered">
          <p>
            <strong>Find the project code in</strong> <a href="https://github.com/pavanas"><i class="fa fa-github"></i></a>
          </p>
            <strong>Coded by @pavanas</strong>
        </div>
      </div>
    </footer>
  </div>
</template>

<script charset="utf-8">
export default {
  data() {
    return {
      meetings: [],
      races: []
    };
  },
  created() {
    this.getInitialData();
  },
  watch: {
    $route: "getInitialData"
  },
  methods: {
    getInitialData: function() {
      // Get meetings data
      let req = $.getJSON("/api/meetings");
      req.done(data => {
        console.log(data);
        this.meetings = data;
      });

      // Get races data
      req = $.getJSON("/api/races");
      req.done(data => {
        console.log(data);
        this.races = data;
      });
    },
    showMeetingRaces: function(event) {
      // check if another element is active
      let elem = document.querySelector('figure.active');
      let meeting_id, is_active = false;

      if(event.target.localName === 'figure') {
        if(elem && elem.id !== event.target.id) {
          elem.classList.toggle('active');
        }
        if (event.target.classList.contains('active')) {
          event.target.classList.toggle('active');
          is_active = false;
        } else {
          event.target.classList.toggle('active');
          meeting_id = event.target.id;
          is_active = true;
        }
      } else {
        if(elem && elem.id !== event.target.parentNode.id) {
          elem.classList.toggle('active');
        }
        if (event.target.parentNode.classList.contains('active')) {
          event.target.parentNode.classList.toggle('active');
          is_active = false;
        } else {
          event.target.parentNode.classList.toggle('active');
          meeting_id = event.target.parentNode.id;
          is_active = true;
        }
      }
      if(is_active) {
        let req = $.getJSON(`api/meetings/${meeting_id}/races`);
        req.done(data => {
          this.races = [];
          console.log(data);
          this.races = data;
        });
      } else {
        this.getInitialData();
      }
    },
    updateRaces: function(event) {
      
      let elem = document.querySelector('.races-content');
      elem.firstChild.remove();
    }
  }
};
</script>