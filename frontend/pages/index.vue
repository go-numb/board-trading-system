<template>
  <section class="section">
    <div class="container">
      <div class="columns">
        <div class="column">
          <p>
            <b-button type="is-success" @click="get" outlined>reload</b-button>
          </p>
          <p v-if="res">{{res}}</p>
          <p v-if="errors">{{errors}}</p>
          <p v-if="board.updated_at != ''" class="has-text-right">Updated: {{board.updated_at}}</p>
          <table class="table is-fullwidth is-hoverable">
            <tbody>
              <tr v-for="(v,i) in board.asks" :key="'ask:'+i">
                <td class="has-text-centered" width="20%">{{v.size}}</td>
                <th class="has-text-centered" width="20%">{{v.price}}</th>
                <td class="has-text-centered" width="20%"></td>
                <th class="has-text-centered" width="20%"></th>
                <td class="has-text-centered" width="20%"></td>
              </tr>
              <tr>
                <td class="has-text-centered"></td>
                <th class="has-text-centered"></th>
                <td class="has-text-centered">{{board.ltp + '-' + board.spread}}</td>
                <th class="has-text-centered"></th>
                <td class="has-text-centered"></td>
              </tr>
              <tr v-for="(v,i) in board.bids" :key="'bid:'+i">
                <td class="has-text-centered"></td>
                <th class="has-text-centered"></th>
                <td class="has-text-centered"></td>
                <th class="has-text-centered">{{v.price}}</th>
                <td class="has-text-centered">{{v.size}}</td>
              </tr>
            </tbody>
          </table>
        </div>
        <div class="column">
          <div class="columns">
            <div class="column field">
              <div class="control has-text-left">
                <b-radio v-model.number="order.type" native-value="0">成行</b-radio>
              </div>
            </div>
            <div class="column field">
              <div class="control has-text-right">
                <b-radio v-model.number="order.type" native-value="1">指値</b-radio>
              </div>
            </div>
          </div>
          <b-field label="Price">
            <b-input v-model.number="order.price"></b-input>
          </b-field>
          <b-field label="Size">
            <b-input v-model.number="order.size"></b-input>
          </b-field>
          <div class="columns">
            <div class="column field">
              <div class="control has-text-left">
                <b-button class="is-danger" @click="sendorder(1)">BUY</b-button>
              </div>
            </div>
            <div class="column field">
              <div class="control has-text-right">
                <b-button class="is-info" @click="sendorder(2)">SELL</b-button>
              </div>
            </div>
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
      // ltp: 0,
      // asks: [],
      // bids: [],
      // updated: "",
      board: [],

      order: {
        type: 0,
        side: 0,
        price: 0,
        size: 0
      },
      oResponse: ""
    };
  },

  mounted: function () {
    this.get()
  },

  methods: {
    get: function() {
      this.$axios
        .$get("http://localhost:8080/api/v1/board")
        .then(res => {
          console.log(res);
          this.res = res;
          this.board = res.data;
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
          if (res.code == 200) {
            // update board
            this.get()
          }
        })
        .catch(err => {
          console.error(err);
          this.oResponse = err;
        });
    }
  }
};
</script>
