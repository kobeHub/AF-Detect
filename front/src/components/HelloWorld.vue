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
            支持上传.mat, .csv格式的心电信号用于分类
          </h3>
        </div>
      </div>
    </section>

    <section class="hero is-medium is-light">
      <div class="hero-body">
        <div class="container">
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
import axios from 'axios'
import JSON from 'JSON'
import qs from 'qs'

export default {
  name: 'HelloWorld',
  data () {
    return {
      primary: 'Atrial Fibrillation Detection Online',
      sub: 'Protect your health in an efficient way',
      nsrdb: nsrdb,
      logo: logo,
      upload_file: null,
      analysis_done: false,
      isLoading: false,
      charData: {
        columns: ['心脏节拍类型', '可能性'],
        rows: [
          {'心脏节拍类型': 'AF节拍', '可能性': 0.875},
          {'心脏节拍类型': '正常节拍', '可能性': 0.025},
          {'心脏节拍类型': '未知', '可能性': 0.1}
        ]
      }
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
      this.analysis_done = false

      // 读取文件内容
      let reader = new FileReader()
      let success = () => {
        this.alertLog({
          title: '系统提示',
          msg: '文件 ' + file.name + ' 上传成功!',
          confirmText: 'Cool!'
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

        this.file_ts = Date.parse(new Date())
        data = JSON.stringify({
          name: e.target.files[0].name,
          ts: this.file_ts,
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
      } else if (this.analysis_done) {
        let msg = {
          title: '系统提示',
          msg: '病情已分析完毕，请重新上传ECG信号分析!',
          confirmText: '确认'
        }
        this.alertLog(msg)
      } else {
        let url = '/api/detect'
        let config = {headers: {'Content-Type': 'application/x-www-form-urlencoded'}}
        let data = qs.stringify({
          ops: 'detect_rencent'
        })

        this.isLoading = true
        axios.post(url, data, config)
          .then(response => {
            console.log(response.data)
          })
          .catch(error => {
            this.alertError(error)
          })

        setTimeout(() => {
          this.isLoading = false
        }, 5 * 1000)
        this.analysis_done = true
      }
    },
    clearFile: function () {
      this.upload_file = null
      this.analysis_done = false
    }
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
