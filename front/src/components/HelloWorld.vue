<template>
  <div class="hello">
    <nav class="navbar" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
        <a class="navbar-item" href="/">
          <img :src="logo" width="28" height="28">
        </a>

        <a role="button" class="navbar-burger burger" aria-label="menu" aria-expanded="True" data-target="navbarBasicExample">
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
        </a>
      </div>

      <div id="navbarBasicExample" class="navbar-menu">
        <div class="navbar-start">
          <a class="navbar-item" href="/">
            Home
          </a>

          <a class="navbar-item" href="https://github.com/kobehub">
            Contact
          </a>
        </div>

        <div class="navbar-end">
          <div class="navbar-item">
            <div class="buttons">
              <a class="button is-primary">
                <strong>注册</strong>
              </a>
              <a class="button is-light">
                登录
              </a>
            </div>
          </div>
        </div>

      </div>
    </nav>

    <section class="hero is-light">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">
            {{ primary }}
          </h1>
          <h2 class="subtitle">
            {{ sub }}
          </h2>
          <img :src="nsrdb">

          <h3 class="is-large">
            支持上传 <strong>*.mat *.hea</strong> 格式的心电信号用于分类
          </h3>
        </div>
      </div>
    </section>

    <section class="hero is-medium is-light">
      <div class="hero-body">
        <div class="container">
          <h2 class="subtitle">病情可视化分析</h2>
          <ve-pie :data="charData"></ve-pie>
          <div class="field">
            <div class="file is-medium is-centered is-boxed is-success has-name">
              <label class="file-label">
                <input class="file-input" type="file" name="ecg_raw" @change="uploadFile">
                  <span class="file-cta">
                    <span class="file-icon">
                       <i class="fas fa-upload"></i>
                    </span>
                    <span class="file-label">
                      点击上传ECG原始信号
                    </span>
                  </span>
                  <span class="file-name">
                    {{ upload_file }}
                  </span>
              </label>
            </div>
          </div>

          <div class="container">
            <FileTable ref="ecgRecord" :pieData="selectUp" @success='updatePie'></FileTable>
          </div>

          <b-button class="button is-primary" @click="analysisConfirm" size="is-medium">AF 诊断</b-button>
          <b-button class="button is-success" @click="clearFile" size="is-medium">清空</b-button>
          <b-loading :is-full-page="true" :active.sync="isLoading" :can-cancel="false">
            <b-icon
                    pack="fas"
                    icon="sync-alt"
                    size="is-large"
                    custom-class="fa-spin">
            </b-icon>
          </b-loading>
        </div>

      </div>
    </section>

    <footer class="footer">
      <div class="content has-text-centered">
        <p>
          <strong>AF-Online</strong> by <a href="https://kobehub.github.io/">Inno Jia</a>. The source code is licensed
          <a href="http://opensource.org/licenses/mit-license.php">MIT</a>.
        </p>
      </div>
    </footer>
  </div>

</template>

<script>
import nsrdb from './../../static/nsrdb.png'
import logo from './../../static/logo.png'
import FileTable from '@/components/FileTable.vue'
import axios from 'axios'
import JSON from 'JSON'

export default {
  name: 'HelloWorld',
  data () {
    return {
      primary: 'Atrial Fibrillation Detection Online',
      sub: 'Protect your health in an efficient way',
      nsrdb: nsrdb,
      logo: logo,
      isLoading: false,
      upload_file: null,
      charData: {
        columns: ['心脏节拍类型', '可能性'],
        rows: [
          {'心脏节拍类型': 'AF节拍', '可能性': 0.875},
          {'心脏节拍类型': '正常节拍', '可能性': 0.025},
          {'心脏节拍类型': '未知', '可能性': 0.1}
        ]
      },
      selectUp: null
    }
  },
  methods: {
    alertLog: function (m) {
      this.$dialog.alert({
        title: m.title,
        message: m.msg,
        confirmText: m.confirmText
      })
    },
    alertError: function (msg) {
      this.$dialog.alert({
        title: '错误',
        message: msg,
        type: 'is-danger',
        hasIcon: true,
        icon: 'times-circle',
        iconPack: 'fa',
        confirmText: '确认'
      })
    },
    uploadFile: function (e) {
      // let formData = new FormData()
      // formData.append('file', e.target.files[0])
      let file = e.target.files[0]
      let data = null // The raw json data
      this.upload_file = e.target.files[0].name
      const TS = new Date()

      // 读取文件内容
      let reader = new FileReader()

      // 文件上传成功，更新文件列表
      let success = () => {
        this.alertLog({
          title: '系统提示',
          msg: '文件 ' + file.name + ' 上传成功!',
          confirmText: 'Cool!'
        })
        this.$refs.ecgRecord.visible = true
        let recordNo = file.name.split('.')[0]
        let recordsInTable = this.$refs.ecgRecord.records
        for (let index in recordsInTable) {
          if (recordsInTable[index].record_name === recordNo) {
            this.$refs.ecgRecord.records[index].files_name.push(file.name)
            return
          }
        }

        // 第一个文件上传
        this.$refs.ecgRecord.records.push({
          record_name: recordNo,
          files_name: [file.name],
          date: TS.toLocaleString(),
          result: null
        })
      }

      let failed = error => {
        this.alertError(error)
      }

      reader.onloadend = function () {
        let url = '/api/upload'

        let arr = new Uint8Array(this.result)
        let arr1 = []
        for (let v in arr) {
          arr1.push(Number(arr[v]))
        }

        // this.file_ts = Date.parse(new Date())
        data = JSON.stringify({
          name: e.target.files[0].name,
          ts: Date.parse(TS),
          content: arr1})

        // application/x-www-form-urlencoded; charset=UTF-8
        let config = {headers: {'Content-Type': 'application/json'}}
        axios.post(url, data, config)
          .then(response => {
            success()
          })
          .catch(error => {
            console.log(error)
            failed(error)
          })
      }

      reader.readAsArrayBuffer(file)
    },
    analysisConfirm: function () {
      if (this.upload_file == null) {
        this.alertError('请上传包含ECG信号的文件!')
      } else if (!this.$refs.ecgRecord.selected) {
        this.alertLog({
          title: '系统提示',
          msg: '请选择列表中的一项进行检测!',
          confirmText: '确定'
        })
      } else if (!this.$refs.ecgRecord.uploadDone()) {
        this.alertLog({
          title: '系统提示',
          msg: '请同时上传.hea .mat文件用于AF检测!',
          confirmText: '确定'
        })
      } else if (this.$refs.ecgRecord.selected.result != null) {
        let msg = {
          title: '系统提示',
          msg: '病情已分析完毕，请选择其他的ECG信号分析!',
          confirmText: '确认'
        }
        this.alertLog(msg)
      } else {
        let url = '/api/detect'
        let config = {headers: {'Content-Type': 'application/json'}}

        // Get selected record from refs.ecgRecord
        let selected = this.$refs.ecgRecord.selected

        let data = JSON.stringify({
          name: selected.record_name
        })

        this.isLoading = true
        axios.post(url, data, config)
          .then(response => {
            console.log(response.data)
            this.charData.rows = [
              {'心脏节拍类型': 'AF节拍', '可能性': response.data.AF},
              {'心脏节拍类型': '正常节拍', '可能性': response.data.NO},
              {'心脏节拍类型': '未知', '可能性': response.data.UN}
            ]
            this.$refs.ecgRecord.setResult(response.data)
            this.isLoading = false
          })
          .catch(error => {
            this.alertError(error)
          })
      }
    },
    clearFile: function () {
      this.upload_file = null
      this.$refs.ecgRecord.clear()
    },
    updatePie: function (result) {
      if (result.result !== null) {
        this.charData.rows = [
          {'心脏节拍类型': 'AF节拍', '可能性': result.result.AF},
          {'心脏节拍类型': '正常节拍', '可能性': result.result.NO},
          {'心脏节拍类型': '未知', '可能性': result.result.UN}
        ]
      }
    }
  },
  components: {
    FileTable
  }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.button {
  margin-right: 60px;
  margin-left: 60px;
  margin-top: 30px;
}
</style>
