<template>
  <div class="row">
    <div
      style="border: 1px solid" :style="infoStyle"
      :class="{'col-md-4': showDetail, 'col-md-12': !showDetail }">
      <div v-for="item of infoItems" :key="item">
        <p>item {{item}}</p><br />
      </div>
      <div style="border: 1px solid" v-pin="{test: 'option1'}">
        <div @click="toggleDetail()"><p>click</p></div>
      </div>
    </div>

    <div>
      <detail ref="Detail"  :style="detailStyle"/>
    </div>
  </div>
</template>

<script>
import Detail from '@/views/detail.vue';

export default {
  name: 'info',
  components: { Detail },
  data () {
    return {
      infoItems: [1, 2, 3],
      showDetail: true,
      detailStyle: { display: 'block' },
      infoStyle: { paddingRight: '30px' }
    };
  },
  methods: {
    toggleDetail () {
      this.showDetail = !this.showDetail;
      this.$log.debug('show details: ', this.showDetail);
      if (this.showDetail) {
        this.detailStyle.display = 'block';
      } else {
        this.detailStyle.display = 'none';
      }
    }
  }
};
</script>
