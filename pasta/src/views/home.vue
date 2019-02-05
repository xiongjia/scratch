<template>
  <div class='home'>
    <Dropdown @on-click="menuClick">
      <Button type="primary">menu<Icon type="ios-arrow-down"></Icon></Button>
      <DropdownMenu slot="list">
        <DropdownItem name="menu1">menu item 1</DropdownItem>
        <DropdownItem name="menu2" disabled>menu item 2</DropdownItem>
        <DropdownItem name="menu3" divided>item 3</DropdownItem>
      </DropdownMenu>
    </Dropdown>

    <ul>
      <li v-for="item in settings" :key="item.id">{{item.name}}<br></li>
    </ul>
  </div>
</template>

<script>
import {mapState, mapGetters, mapActions} from 'vuex';

const mixin = {
  created () {
    this.$log.debug('mixin created');
  },
  methods: {
    getMessage () {
      return 'test';
    }
  }
};

export default {
  name: 'home',
  mixins: [mixin],
  components: { },
  data () {
    return {
      users: [{name: 'test1'}]
    };
  },
  created () {
    this.$log.debug('view(home) is created');
  },
  computed: {
    ...mapState('configData', ['settings'])
  },
  methods: {
    ...mapActions('configData', ['loadConfig']),
    ...mapGetters('configData', ['getItem', 'getData']),
    menuClick (name) {
      this.$log.debug('menu clicked', name);
      this.$log.debug('message is ', this.getMessage());
      this.loadConfig({item: 'dispatch'});
      console.log(this.getData());
      console.log('item ', this.getItem()(0));
    }
  }
};
</script>
