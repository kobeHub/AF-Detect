<template>
  <section v-if="visible">
    <b-tabs>
      <b-tab-item label="Table">
        <b-table :data="records"
                    :columns="columns"
                    :selected.sync="selected"
                    focusable>
        </b-table>
      </b-tab-item>

      <b-tab-item label="Selected">
          <pre>{{ selected }}</pre>
      </b-tab-item>
    </b-tabs>
  </section>
</template>

<script>
export default {
  name: 'FileTable',
  props: ['pieData'],
  data () {
    return {
      records: [],
      selected: null,
      visible: false,
      columns: [
        {
          field: 'record_name',
          label: 'ECG信号'
        },
        {
          field: 'files_name',
          label: '包含文件'
        },
        {
          field: 'date',
          label: '日期',
          centered: true
        },
        {
          field: 'result',
          label: '检测结果'
        }
      ]
    }
  },
  methods: {
    getRecord: function (reno) {
      for (let index in this.records) {
        if (this.records[index].record_name === reno) {
          return this.records[index]
        }
      }
    },
    setResult: function (res) {
      this.selected.result = res
    },
    uploadDone: function () {
      return this.selected.files_name.length > 1
    },
    clear: function () {
      this.records = []
      this.visible = false
    },
    existsEle: function () {
      return this.records !== []
    }
  },
  watch: {
    selected: {
      handler () {
        this.$emit('success', this.selected)
      }
    }
  }
}
</script>

<style scoped>

</style>
