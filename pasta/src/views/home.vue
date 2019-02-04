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
  </div>
</template>

<script>
export default {
  name: 'home',
  components: { },
  data () {
    return {
      users: [{name: 'test1'}]
    };
  },
  created () {
    this.$log.debug('view(home) is created');
  },
  methods: {
    menuClick (name) {
      console.log('menu clicked', name);
    }
  }
};
</script>
