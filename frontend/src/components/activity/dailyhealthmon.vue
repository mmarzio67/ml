<template>
  <div id="Activity">
    <div class="activity-form">
      
      <div class="submit">
          <button type="submit" @click="updateLastHealthValues()">Last health entry</button>
        </div>
      <form @submit.prevent="onSubmit">
        <div class="input">
          <label for="weight">Weight:</label>
          <input type="number" step="0.1" id="weight" v-model="weight" />
        </div>
        <div class="input">
          <label for="bfi">Body fat:</label>
          <input type="number" step="0.1" id="bfi" v-model="bfi" />
        </div>
        <div class="input">
          <label for="imc">IMC:</label>
          <input type="number" step="0.1" id="imc" v-model="imc" />
        </div>
        <div class="input">
          <label for="waist">Waist:</label>
          <input type="number" id="waist" v-model.number="waist" />
        </div>
        <div class="input">
          <label for="spo2">SPO2:</label>
          <input
            type="number"
            id="spo2"
            min="50"
            max="100"
            v-model.number="spo2"
          />
        </div>
        <div class="input">
          <label for="breathrest">Breath at rest:</label>
          <input type="number" id="breathrest" v-model.number="breathrest" />
        </div>
        <div class="input">
          <label for="breathactive">Breath active:</label>
          <input
            type="number"
            id="breathactive"
            v-model.number="breathactive"
          />
        </div>
        <div class="input">
          <label for="hslept">Hours of Sleep:</label>
          <input type="time" v-model="hslept" />
        </div>
        <div class="input">
          <label for="pulserest">Pulse at rest:</label>
          <input
            type="number"
            id="pulserest"
            min="30"
            max="180"
            v-model.number="pulserest"
          />
        </div>
        <div class="input">
          <label for="pulseactive">Pulse active:</label>
          <input
            type="number"
            id="pulseactive"
            min="30"
            max="180"
            v-model.number="pulseactive"
          />
        </div>
        <div class="input">
          <label for="stress">Stress:</label>
          <input type="number" id="stress" v-model.number="stress" />
        </div>
        <div class="input">
          <label for="bodybattU">Body Battery UP (+)</label>
          <input type="number" id="bodybatt" v-model.number="bodybattU" />
        </div>
        <div class="input">
          <label for="bodybattD">Body Battery DOWN (-)</label>
          <input type="number" id="bodybatt" v-model.number="bodybattD" />
        </div>

        <div class="input">
          <label for="steps">Steps</label>
          <input type="number" id="steps" v-model.number="steps" />
        </div>
        <div class="submit">
          <button type="submit">Submit</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      weight: null,
      bfi: null,
      imc: null,
      waist: null,
      spo2: null,
      breathrest: null,
      breathactive: null,
      hslept: null,
      pulserest: null,
      pulseactive: null,
      stress: null,
      bodybattU: null,
      bodybattD: null,
      steps: null,
    };
  },
  computed: {
    hourslept() {
      var hoursMinutes = this.hslept.split(/[.:]/);
      var hours = parseInt(hoursMinutes[0], 10);
      var minutes = hoursMinutes[1] ? parseInt(hoursMinutes[1], 10) : 0;
      return hours + minutes / 60;
    },
  },
  created() {
    this.$store.dispatch("fetchUser");
    this.$store.dispatch("lastHealthEntry");
    
  },
  methods: {
    onSubmit() {
      const formData = {
        weight: this.weight,
        bfi: this.bfi,
        imc: this.imc,
        waist: this.waist,
        spo2: this.spo2,
        breathrest: this.breathrest,
        breathactive: this.breathactive,
        timeslept: this.hourslept,
        pulserest: this.pulserest,
        pulseactive: this.pulseactive,
        stress: this.stress,
        bodybattU: this.bodybattU,
        bodybattD: this.bodybattD,
        steps: this.steps,
      };
      console.log("onSubmit:", formData);
      this.$store.dispatch("dailyhealthmon", formData);
    },
    updateLastHealthValues(){ 
    const lhr = this.$store.getters["getLastHealthEntry"];
    console.log(lhr)
    this.weight=lhr.weight,
    this.bfi=lhr.bfi,
    this.imc=lhr.imc,
    this.waist=lhr.waist,
    this.spo2=lhr.spo2,
    this.breathrest=lhr.breathrest,
    this.breathactive=lhr.breathactive,
    this.hslept=lhr.hslept,
    this.pulserest=lhr.pulserest,
    this.pulseactive=lhr.pulseactive,
    this.stress=lhr.stress,
    this.bodybattU=lhr.bodybattU,
    this.bodybattD=lhr.bodybattD,
    this.steps=lhr.steps
    }
  },
  sortByDate(){

  }
};
</script>

<style scoped>
.activity-form {
  width: 400px;
  margin: 30px auto;
  border: 1px solid #eee;
  padding: 20px;
  box-shadow: 0 2px 3px #ccc;
}

.input,
.input-group {
  margin: 10px auto;
}

.input label {
  display: block;
  color: #4e4e4e;
  margin-bottom: 6px;
}

.input.inline label {
  display: inline;
}

.input input {
  font: inherit;
  width: 100%;
  padding: 6px 12px;
  box-sizing: border-box;
  border: 1px solid #ccc;
}

.input.inline input {
  width: auto;
}

.input input:focus {
  outline: none;
  border: 1px solid #521751;
  background-color: #eee;
}

.input select {
  border: 1px solid #ccc;
  font: inherit;
}

.hobbies button {
  border: 1px solid #521751;
  background: #521751;
  color: white;
  padding: 6px;
  font: inherit;
  cursor: pointer;
}

.hobbies button:hover,
.hobbies button:active {
  background-color: #8d4288;
}

.hobbies input {
  width: 90%;
}

.submit button {
  border: 1px solid #521751;
  color: #521751;
  padding: 10px 20px;
  font: inherit;
  cursor: pointer;
}

.submit button:hover,
.submit button:active {
  background-color: #521751;
  color: white;
}

.submit button[disabled],
.submit button[disabled]:hover,
.submit button[disabled]:active {
  border: 1px solid #ccc;
  background-color: transparent;
  color: #ccc;
  cursor: not-allowed;
}
</style>
