<template>
  <div>
    <div style="border: 1px solid" :style="infoStyle">
      <p>info view</p><br />
      <div @click="toggleDetail()"><p>click</p></div>
    </div>

    <detail ref="Detail"  :style="detailStyle"/>
  </div>
</template>

<script>
import Detail from '@/views/detail.vue';

export default {
  name: 'info',
  components: { Detail },
  data () {
    return {
      showDetail: false,
      detailStyle: { display: 'none' },
      infoStyle: { paddingRight: '0' }
    };
  },
  methods: {
    toggleDetail () {
      this.showDetail = !this.showDetail;
      this.$log.debug('show details: ', this.showDetail);
      if (this.showDetail) {
        this.detailStyle.display = 'block';
        this.infoStyle.paddingRight = '300px';
      } else {
        this.detailStyle.display = 'none';
        this.infoStyle.paddingRight = '0';
      }
    }
  }
};
</script>
