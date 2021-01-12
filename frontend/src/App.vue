<template>
  <v-app id="inspire">
    <v-main>
      <v-container fluid>
        <v-layout justify-center align-center class="main-layout">
          <main>
            <header>
              <h1>
                VirusTotal File Checker
              </h1>
            </header>

            <!-- display file and upload information -->
            <div class="file-info-upload" v-if="uploading">
              <p>
                Filename:
                <span class="filename">{{ upload_filename }}</span>
              </p>
              <hr />
              <div class="uploading">
                <p>{{ upload_status }}</p>
                <v-progress-linear
                  indeterminate
                  color="green"
                  v-if="!done_upload"
                ></v-progress-linear>
              </div>
            </div>
            <!-- results button actons -->
            <div class="results-main" v-if="done_upload">
              <Result
                :data="result"
                :upload_filename="upload_filename"
                @reset-new="ResetNewSession"
              />
            </div>

            <!-- main form -->
            <div class="main-form" v-if="form_upload">
              <p>Please select a file to upload and check...</p>
              <v-btn x-large block primary @click="SelectFile">
                Choose a file...
              </v-btn>
            </div>

            <!-- footer -->
            <footer>
              <v-dialog v-model="dialog" persistent max-width="500">
                <template v-slot:activator="{ on, attrs }">
                  <button v-bind="attrs" v-on="on">
                    <v-icon x-large class="settings-btn">mdi-cogs</v-icon>
                  </button>
                </template>
                <v-card>
                  <v-card-title class="headline">
                    Your VirusTotal API Key
                  </v-card-title>
                  <v-container>
                    <v-col cols="12">
                      <v-text-field
                        label="Api Key*"
                        required
                        v-model="form_api"
                      ></v-text-field>
                    </v-col>
                  </v-container>
                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn
                      color="gray darken-1"
                      text
                      @click="dialog = form_api.length > 0 ? false : true"
                      outlined
                    >
                      Close
                    </v-btn>
                    <v-btn
                      color="green darken-1"
                      text
                      @click="SaveAPI"
                      outlined
                    >
                      Save
                    </v-btn>
                  </v-card-actions>
                </v-card>
              </v-dialog>
            </footer>
          </main>
        </v-layout>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import Result from './components/Results.vue'
import Wails from '@wailsapp/runtime'

export default {
  data: () => ({
    dialog: false,
    form_api: '',
    form_upload: true,
    uploading: false,
    upload_filename: '',
    upload_status: 'Uploading...',
    done_upload: false,
    result: {},
  }),
  watch: {
    upload_filename: {
      handler(upload_filename) {
        if (upload_filename === '') {
          return ''
        }
        this.uploading = true
        this.form_upload = false

        // run the uploader
        this.UploadFile()
      },
    },
    result: {
      handler(result) {
        if (Object.keys(result).length > 0) {
          // upload status
          this.upload_status = 'Done...'

          // set to done
          this.done_upload = true
        }
      },
    },
  },
  methods: {
    upload() {
      this.upload_filename = '/home/theboringdude/elyca'
    },
    // create a new session and reset all vars (except the api)
    ResetNewSession() {
      this.form_upload = true
      this.uploading = false
      this.upload_filename = ''
      this.upload_status = 'Uploading...'
      this.done_upload = false
      this.result = {}
    },
    // handles selecting file to be uploaded
    SelectFile() {
      window.backend.File.SelectFileUpload().then((file) => {
        if (file !== '') {
          this.upload_filename = file
        }
      })
    },
    // handles loading the api file to the app
    LoadAPI() {
      window.backend.API.Load()
        .then((res) => {
          const data = JSON.parse(res)
          this.form_api = data.api

          // check the loaded api
          this.CheckAPIFile()
        })
        .catch((e) => console.error(e))
    },
    // handles saving a new api
    SaveAPI() {
      if (this.form_api.length > 0 || this.form_api !== undefined) {
        const NewAPI = {
          api: this.form_api,
        }

        // save new api
        window.backend.API.Save(JSON.stringify(NewAPI, null, 2))
          .then()
          .catch((e) => console.error(e))
      }
    },
    // handles checkingif the api is valid or not
    CheckAPIFile() {
      // check if the api exists in the file
      if (this.form_api === undefined) {
        this.dialog = true // show if none
      } else {
        this.dialog = false
      }
    },
    // upload files
    UploadFile() {
      // call the backend upload function
      // NOTE: the app will not work if the frontend will
      // be the one to try and upload it
      window.backend.File.FileUpload(this.form_api)
        .then((res) => {
          // convert response to json
          const data = JSON.parse(res)
          const resp = JSON.parse(data.response)

          // the statuscode should be '200'
          // others are errors
          if (data.statusCode == 200) {
            this.upload_status = `Retrieving results...`

            // get another request
            this.RequestAnalysis(resp.data.id)
          }
        })
        .catch((e) => console.error(e))
    },
    // get again the request analysis after upload
    RequestAnalysis(analysisId) {
      // using axios is not working, so I used the backend
      // to send a request and retrieve back the analysis results
      window.backend.File.GetAnalysisFromID(analysisId, this.form_api)
        .then((res) => {
          this.result = JSON.parse(res)
        })
        .catch((e) => console.error(e))
    },
  },
  mounted() {
    Wails.Events.On('datamodified', () => {
      this.LoadAPI()
    })

    // load on startup
    this.LoadAPI()
  },
  components: {
    Result,
  },
}
</script>

<style>
.file-info-upload {
  padding: 2em;
  margin-top: 12px;
  border: 1px solid #e0e0e0;
}
.file-info-upload p {
  font-size: 1.2em;
}
.file-info-upload .filename {
  font-weight: bold;
  text-decoration: underline;
  color: #4caf50;
}
.file-info-upload .uploading {
  margin-top: 10px;
}
.file-info-upload .uploading p {
  font-size: 1em;
}
.file-info-upload hr {
  border: 1px solid #e0e0e0;
}
.results-main {
  text-align: center;
  margin-top: 14px;
}
.results-main button {
  margin: 0 1px 0 1px;
}
.main-layout {
  height: 550px;
  width: 900px;
  margin: 0px auto;
  position: relative;
}
header h1 {
  font-size: 4em;
  letter-spacing: 2px;
  text-align: center;
}
.main-form {
  margin-top: 10px;
}
footer button {
  position: absolute;
  bottom: 0px;
  right: 0px;
}
.settings-btn {
  height: 64px;
  background: #e0e0e0;
  width: 64px;
}
.settings-btn:hover {
  background: #bdbdbd;
}
</style>
