<template lang="pug">
div(class="home")
  Dropdown(@on-click="menuClick")
    Button(type="primary") menu
      Icon(type="ios-arrow-down")
    DropdownMenu(slot="list")
      DropdownItem(name="menu1") menu item 1
      DropdownItem(name="menu2" disabled) menu item 2
      DropdownItem(name="menu3" divided) menu item 3

  ul
    li(v-for="item in settings" :key="item.id") {{item.name}}

  Items(:testItems="testItems")

  DatePicker(type="datetime" :options="datePickerOpts")

  Modal(v-model="modalTest" :closable="false"
    title="Common Modal dialog box title")
    p() Content of dialog
    p() Content of dialog

  Table(:columns="columns" :data="tableData")
</template>

<script>
import {mapState, mapGetters, mapActions} from 'vuex';
import Items from '@/components/items';
import expandRow from './table-expand.vue';

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
  components: {Items, expandRow},
  data () {
    return {
      modalTest: false,
      users: [{name: 'test1'}],
      testItems: {
        k1: 'value1',
        k2: { k3: 'value3' }
      },
      datePickerOpts: {
        disabledDate (date) {
          console.log('date', date);
          const now = Date.now();
          return now < date.getTime();
        }
      },
      columns: [{
        type: 'expand',
        width: 50,
        render: (h, params) => h(expandRow, { props: { row: params.row } })
      }, {
        title: 'Name', key: 'name'
      }, {
        title: 'Age', key: 'age'
      }, {
        title: 'Address', key: 'address'
      }],
      tableData: [{
        name: 'John Brown',
        age: 18,
        address: 'New York No. 1 Lake Park',
        job: 'Data engineer',
        interest: 'badminton',
        birthday: '1991-05-14',
        book: 'Steve Jobs',
        movie: 'The Prestige',
        music: 'I Cry'
      }]
    };
  },
  created () {
    this.$log.debug('view(home) is created');
  },
  computed: {
    ...mapState('configData', ['settings']),
    ...mapGetters('configData', ['getData2'])
  },
  methods: {
    ...mapActions('configData', ['loadConfig']),
    ...mapGetters('configData', ['getItem', 'getData']),
    menuClick (name) {
      this.$log.debug('menu clicked', name);
      this.$log.debug('message is ', this.getMessage());
      this.loadConfig({item: 'dispatch'});
      console.log('getdata', this.getData());
      console.log('getdata2', this.getData2);
      console.log('item ', this.getItem()(0));

      if (name === 'menu1') {
        this.modalTest = true;
      }
    }
  }
};
</script>
