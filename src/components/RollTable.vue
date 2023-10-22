<template>
  <LuckyWheel
      ref="LuckyWheel"
      width="700px"
      height="700px"
      :prizes="prizes"
      :default-style="defaultStyle"
      :blocks="blocks"
      :buttons="buttons"
      @start="startCallBack"
      @end="endCallBack"
  />
</template>

<script>
export default {
  props: {
    prizeList: Array
  },

  data() {
    return {
      prizes: [],
      defaultStyle: {
        fontColor: '#d64737',
        fontSize: '14px'
      },
      blocks: [
        {padding: '20px', background: '#d64737'}
      ],
      buttons: [
        {radius: '100px', background: '#d64737'},
        {radius: '80px', background: '#ff7402', pointer: true},
        {
          radius: '60px',
          background: '#fff',
          fonts: [{text: 'start',fontSize: '40px', top: '-25px'}]
        }
      ],
    }
  },
  mounted() {
    this.getPrizesList()
  },
  methods: {
    getPrizesList() {
      const prizes = []
      let data = this.prizeList
      data.forEach((item, index) => {
        prizes.push({
          title: item.name,
          background: index % 2 ? '#f9e3bb' : '#f8d384',
          fonts: [{text: item.name, top: '10%',fontSize: '20px'}],
          imgs: [{ src: item.src, width: '45%', top: '20%' }],
        })
      })
      this.prizes = prizes
    },
    startCallBack() {
      this.$refs.LuckyWheel.play()
      setTimeout(() => {
        this.$refs.LuckyWheel.stop(Math.random() * 8 >> 0)
      }, 3000)
    },
    endCallBack(prize) {
      // alert(`恭喜你获得${prize.title}`)
      console.log(prize.title, prize.index)
    },
  }
}
</script>
