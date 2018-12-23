<template>
  <div>
    <p>detail view</p>
    <div v-for="item of items" :key="item">
      <p>data {{item}}</p><br/>
    </div>
  </div>
</template>

<script>
import _ from 'lodash';

export default {
  name: 'detail',
  data () {
    return {
      items: _.range(1, 20)
    };
  }
};
</script>
