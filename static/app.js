const vm = new Vue({
  data() {
    return {
      count: 0,
      isAccessing: false,
      socket: null,
    }
  },
  methods: {
    countUp() {
      let app = this
      axios.get("/app/countUp")
        .then(res => {
          app.count = res.data.count
        })
    },
    clear() {
      let app = this
      axios.get("/app/clear")
        .then(res => {
          app.count = res.data.count
        })
    },
  },
  computed: {
    //NONE
  },
  el: "#app"
})
