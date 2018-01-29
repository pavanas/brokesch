<template>
	<div>
		<section class="hero is-info is-medium">
		  <div class="hero-body">
		    <div class="container has-text-centered">
		      <h1 class="title">
		        Next 5
		      </h1>
		      <h2 class="subtitle">
		        <strong>Visualizing the last 5 races for easier betting!</strong>
		      </h2>
		    </div>
		  </div>
		</section>
		<section class="section is-medium">
		  <div class="tile is-ansector">
		    <div class="tile is-4 is-parent raceinfo">
		      <article class="tile is-child">
		        <p class="title">Race info</p>
		        <p>Race number: #{{ race[0].number }}</p>
		        <p>Category: {{ race[0].m_type }}</p>
		        <p>Closing in: <countdown :deadline="[race[0].ends_at]">1h:22m:33s</countdown></p>
		      </article>
		    </div>
		    <div class="tile is-6 is-parent">
		      <article class="tile is-child">
		        <div class="level is-competitor" v-for="competitor in race">
		          <div class="level-item"></div>
		          <div class="level-item">
		          	<div>
		            	<p class="heading is-spaced">Competitor</p>
		           		<p class="title">#{{ competitor.number_c }}</p>
		           	</div>
		          </div>
		          <div class="level-item">
		            <div>
		            	<p class="heading is-spaced">Row</p>
		            	<p class="title">{{ competitor.position }}</p>
		            </div>
		          </div>
		        </div>
		      </article>
		    </div>
		  </div>
		</section>
		<footer class="footer">
		  <div class="container">
		    <div class="content has-text-centered">
		      <p>
		        <strong>Find the project code in</strong> <a href="https://github.com/pavanas/brokesch"><i class="fa fa-github"></i></a>
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
      race: []
    };
  },
  created() {
    this.getRaceInfo();
  },
  watch: {
    $route: "getRaceInfo"
  },
  methods: {
    getRaceInfo: function() {
      const id = this.$route.params.id;
      let req = $.getJSON(`/api/races/${id}`);
      req.done(data => {
      	console.log(data);
        this.race = data;
      });
    }
  }
};
</script>