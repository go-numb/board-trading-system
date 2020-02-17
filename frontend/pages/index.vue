<template>
  <section class="section">
    <div class="container">
      <div class="columns">
        <div class="column">
          <p>
            <b-button type="is-success" @click="get" outlined>Get board</b-button>
          </p>
          <p v-if="res">{{res}}</p>
          <p v-if="errors">{{errors}}</p>
          <table class="table is-fullwidth is-hoverable">
            <tbody>
              <tr v-for="(v,i) in asks" :key="'ask:'+i">
                <td>{{v.size}}</td>
                <th>{{v.price}}</th>
                <td></td>
                <th></th>
                <td></td>
              </tr>
              <tr>
                <td></td>
                <th></th>
                <td>{{ltp + '-' + spread}}</td>
                <th></th>
                <td></td>
              </tr>
              <tr v-for="(v,i) in bids" :key="'bid:'+i">
                <td></td>
                <th></th>
                <td></td>
                <th>{{v.price}}</th>
                <td>{{v.size}}</td>
              </tr>
            </tbody>
          </table>
        </div>
        <div class="column">
          <div class="block">
            <b-radio v-model.number="order.type" native-value="0">成行</b-radio>
            <b-radio v-model.number="order.type" native-value="1">指値</b-radio>
          </div>
          <b-field label="Price">
            <b-input v-model.number="order.price"></b-input>
          </b-field>
          <b-field label="Size">
            <b-input v-model.number="order.size"></b-input>
          </b-field>
          <div class="block">
            <b-button class="is-danger" @click="sendorder(1)">BUY</b-button>
            <b-button class="is-info" @click="sendorder(2)">SELL</b-button>
          </div>
          <p v-if="oResponse">{{oResponse}}</p>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: "HomePage",

  data() {
    return {
      errors: "",
      ltp: 0,
      asks: [],
      bids: [],

      order: {
        type: 0,
        side: 0,
        price: 0,
        size: 0
      },
      oResponse: ""
    };
  },

  methods: {
    get: function() {
      this.$axios
        .$get("http://localhost:8080/api/v1/board")
        .then(res => {
          console.log(res);
          this.res = res;
          this.ltp = res.data.ltp;
          this.spread = res.data.spread;
          this.asks = res.data.asks;
          this.bids = res.data.bids;
        })
        .catch(err => {
          console.error(err);
          this.errors = err;
        });
    },
    sendorder: function(side) {
      // set order side
      this.order.side = side;
      this.$axios
        .$post("http://localhost:8080/api/v1/private/order", this.order, {
          headers: { "Content-Type": "application/json" }
        })
        .then(res => {
          console.log(res);
          this.oResponse = res.status;
        })
        .catch(err => {
          console.error(err);
          this.oResponse = err;
        });
    }
  }
};
</script>
