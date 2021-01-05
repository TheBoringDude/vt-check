<template>
  <!-- full screen results dialog -->
  <v-dialog
    v-model="show_results"
    fullscreen
    hide-overlay
    transition="dialog-bottom-transition"
  >
    <template v-slot:activator="{ on, attrs }">
      <v-btn color="secondary">
        Check Another File
      </v-btn>
      <v-btn color="primary" dark v-bind="attrs" v-on="on">
        Show Results
      </v-btn>
    </template>
    <v-card>
      <v-toolbar dark color="primary">
        <v-btn icon dark @click="show_results = hideResults">
          <v-icon>mdi-close</v-icon>
        </v-btn>
        <v-toolbar-title>Results: {{ upload_filename }}</v-toolbar-title>
        <!-- <v-spacer></v-spacer>
        <v-toolbar-items>
          <v-btn dark text @click="show_results = false">Save</v-btn>
        </v-toolbar-items> -->
      </v-toolbar>
      <v-list three-line subheader>
        <v-subheader>General</v-subheader>
        <v-container>
          <v-row dense>
            <v-col
              cols="4"
              v-for="result in data.data.attributes.results"
              v-bind:key="result.engine_name"
            >
              <v-card>
                <v-list-item three-line>
                  <v-list-item-content>
                    <div class="overline mb-2">{{ result.engine_name }}</div>
                    <v-list-item-title class="headline mb-1">
                      {{ result.category }}
                    </v-list-item-title>
                    <v-list-item-subtitle class="result-subtitle">
                      {{ result.result }}
                    </v-list-item-subtitle>
                  </v-list-item-content>

                  <v-list-item-avatar
                    tile
                    size="60"
                    color="grey"
                  ></v-list-item-avatar>
                </v-list-item>
              </v-card>
            </v-col>
          </v-row>
        </v-container>
      </v-list>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: 'Result',
  computed: {
    hideResults() {
      return false
    },
  },
  data() {
    return {
      show_results: false,
      data: {
        data: {
          attributes: {
            date: 1475048277,
            results: {
              ALYac: {
                category: 'undetected',
                engine_name: 'ALYac',
                engine_update: '20160928',
                engine_version: '1.0.1.9',
                method: 'blacklist',
                result: null,
              },
              AVG: {
                category: 'undetected',
                engine_name: 'AVG',
                engine_update: '20160928',
                engine_version: '16.0.0.4656',
                method: 'blacklist',
                result: null,
              },
              AVware: {
                category: 'undetected',
                engine_name: 'AVware',
                engine_update: '20160928',
                engine_version: '1.5.0.42',
                method: 'blacklist',
                result: null,
              },
              'Ad-Aware': {
                category: 'undetected',
                engine_name: 'Ad-Aware',
                engine_update: '20160928',
                engine_version: '3.0.3.794',
                method: 'blacklist',
                result: null,
              },
              AegisLab: {
                category: 'undetected',
                engine_name: 'AegisLab',
                engine_update: '20160928',
                engine_version: '4.2',
                method: 'blacklist',
                result: null,
              },
              'AhnLab-V3': {
                category: 'undetected',
                engine_name: 'AhnLab-V3',
                engine_update: '20160928',
                engine_version: '3.7.5.15490',
                method: 'blacklist',
                result: null,
              },
              Alibaba: {
                category: 'undetected',
                engine_name: 'Alibaba',
                engine_update: '20160928',
                engine_version: '1.0',
                method: 'blacklist',
                result: null,
              },
              'Antiy-AVL': {
                category: 'undetected',
                engine_name: 'Antiy-AVL',
                engine_update: '20160928',
                engine_version: '1.0.0.1',
                method: 'blacklist',
                result: null,
              },
              Arcabit: {
                category: 'undetected',
                engine_name: 'Arcabit',
                engine_update: '20160928',
                engine_version: '1.0.0.774',
                method: 'blacklist',
                result: null,
              },
              Avast: {
                category: 'undetected',
                engine_name: 'Avast',
                engine_update: '20160928',
                engine_version: '8.0.1489.320',
                method: 'blacklist',
                result: null,
              },
              Avira: {
                category: 'undetected',
                engine_name: 'Avira',
                engine_update: '20160928',
                engine_version: '8.3.3.4',
                method: 'blacklist',
                result: null,
              },
              Baidu: {
                category: 'undetected',
                engine_name: 'Baidu',
                engine_update: '20160928',
                engine_version: '1.0.0.2',
                method: 'blacklist',
                result: null,
              },
              BitDefender: {
                category: 'undetected',
                engine_name: 'BitDefender',
                engine_update: '20160928',
                engine_version: '7.2',
                method: 'blacklist',
                result: null,
              },
              Bkav: {
                category: 'undetected',
                engine_name: 'Bkav',
                engine_update: '20160927',
                engine_version: '1.3.0.8383',
                method: 'blacklist',
                result: null,
              },
              'CAT-QuickHeal': {
                category: 'undetected',
                engine_name: 'CAT-QuickHeal',
                engine_update: '20160928',
                engine_version: '14.00',
                method: 'blacklist',
                result: null,
              },
              CMC: {
                category: 'undetected',
                engine_name: 'CMC',
                engine_update: '20160921',
                engine_version: '1.1.0.977',
                method: 'blacklist',
                result: null,
              },
              ClamAV: {
                category: 'undetected',
                engine_name: 'ClamAV',
                engine_update: '20160928',
                engine_version: '0.98.5.0',
                method: 'blacklist',
                result: null,
              },
              Comodo: {
                category: 'undetected',
                engine_name: 'Comodo',
                engine_update: '20160928',
                engine_version: '25841',
                method: 'blacklist',
                result: null,
              },
              Cyren: {
                category: 'undetected',
                engine_name: 'Cyren',
                engine_update: '20160928',
                engine_version: '5.4.16.7',
                method: 'blacklist',
                result: null,
              },
              DrWeb: {
                category: 'undetected',
                engine_name: 'DrWeb',
                engine_update: '20160928',
                engine_version: '7.0.23.8290',
                method: 'blacklist',
                result: null,
              },
              'ESET-NOD32': {
                category: 'undetected',
                engine_name: 'ESET-NOD32',
                engine_update: '20160928',
                engine_version: '14190',
                method: 'blacklist',
                result: null,
              },
              Emsisoft: {
                category: 'undetected',
                engine_name: 'Emsisoft',
                engine_update: '20160928',
                engine_version: '3.5.0.658',
                method: 'blacklist',
                result: null,
              },
              'F-Prot': {
                category: 'undetected',
                engine_name: 'F-Prot',
                engine_update: '20160926',
                engine_version: '4.7.1.166',
                method: 'blacklist',
                result: null,
              },
              'F-Secure': {
                category: 'undetected',
                engine_name: 'F-Secure',
                engine_update: '20160928',
                engine_version: '11.0.19100.45',
                method: 'blacklist',
                result: null,
              },
              Fortinet: {
                category: 'undetected',
                engine_name: 'Fortinet',
                engine_update: '20160928',
                engine_version: '5.4.233.0',
                method: 'blacklist',
                result: null,
              },
              GData: {
                category: 'undetected',
                engine_name: 'GData',
                engine_update: '20160928',
                engine_version: '25',
                method: 'blacklist',
                result: null,
              },
              Ikarus: {
                category: 'undetected',
                engine_name: 'Ikarus',
                engine_update: '20160927',
                engine_version: 'T3.2.1.6.0',
                method: 'blacklist',
                result: null,
              },
              Jiangmin: {
                category: 'undetected',
                engine_name: 'Jiangmin',
                engine_update: '20160928',
                engine_version: '16.0.100',
                method: 'blacklist',
                result: null,
              },
              K7AntiVirus: {
                category: 'undetected',
                engine_name: 'K7AntiVirus',
                engine_update: '20160928',
                engine_version: '9.241.21013',
                method: 'blacklist',
                result: null,
              },
              K7GW: {
                category: 'undetected',
                engine_name: 'K7GW',
                engine_update: '20160928',
                engine_version: '9.241.21013',
                method: 'blacklist',
                result: null,
              },
              Kaspersky: {
                category: 'undetected',
                engine_name: 'Kaspersky',
                engine_update: '20160928',
                engine_version: '15.0.1.13',
                method: 'blacklist',
                result: null,
              },
              Kingsoft: {
                category: 'undetected',
                engine_name: 'Kingsoft',
                engine_update: '20160928',
                engine_version: '2013.8.14.323',
                method: 'blacklist',
                result: null,
              },
              Malwarebytes: {
                category: 'undetected',
                engine_name: 'Malwarebytes',
                engine_update: '20160928',
                engine_version: '2.1.1.1115',
                method: 'blacklist',
                result: null,
              },
              McAfee: {
                category: 'undetected',
                engine_name: 'McAfee',
                engine_update: '20160928',
                engine_version: '6.0.6.653',
                method: 'blacklist',
                result: null,
              },
              'McAfee-GW-Edition': {
                category: 'undetected',
                engine_name: 'McAfee-GW-Edition',
                engine_update: '20160927',
                engine_version: 'v2015',
                method: 'blacklist',
                result: null,
              },
              'MicroWorld-eScan': {
                category: 'undetected',
                engine_name: 'MicroWorld-eScan',
                engine_update: '20160928',
                engine_version: '12.0.250.0',
                method: 'blacklist',
                result: null,
              },
              Microsoft: {
                category: 'undetected',
                engine_name: 'Microsoft',
                engine_update: '20160928',
                engine_version: '1.1.13103.0',
                method: 'blacklist',
                result: null,
              },
              'NANO-Antivirus': {
                category: 'undetected',
                engine_name: 'NANO-Antivirus',
                engine_update: '20160927',
                engine_version: '1.0.38.11617',
                method: 'blacklist',
                result: null,
              },
              Panda: {
                category: 'undetected',
                engine_name: 'Panda',
                engine_update: '20160927',
                engine_version: '4.6.4.2',
                method: 'blacklist',
                result: null,
              },
              'Qihoo-360': {
                category: 'undetected',
                engine_name: 'Qihoo-360',
                engine_update: '20160928',
                engine_version: '1.0.0.1120',
                method: 'blacklist',
                result: null,
              },
              Rising: {
                category: 'undetected',
                engine_name: 'Rising',
                engine_update: '20160928',
                engine_version: '28.0.0.1',
                method: 'blacklist',
                result: null,
              },
              SUPERAntiSpyware: {
                category: 'undetected',
                engine_name: 'SUPERAntiSpyware',
                engine_update: '20160928',
                engine_version: '5.6.0.1032',
                method: 'blacklist',
                result: null,
              },
              Sophos: {
                category: 'undetected',
                engine_name: 'Sophos',
                engine_update: '20160928',
                engine_version: '4.98.0',
                method: 'blacklist',
                result: null,
              },
              Symantec: {
                category: 'undetected',
                engine_name: 'Symantec',
                engine_update: '20160928',
                engine_version: '20151.1.1.4',
                method: 'blacklist',
                result: null,
              },
              Tencent: {
                category: 'undetected',
                engine_name: 'Tencent',
                engine_update: '20160928',
                engine_version: '1.0.0.1',
                method: 'blacklist',
                result: null,
              },
              TheHacker: {
                category: 'undetected',
                engine_name: 'TheHacker',
                engine_update: '20160927',
                engine_version: '6.8.0.5.1080',
                method: 'blacklist',
                result: null,
              },
              TrendMicro: {
                category: 'undetected',
                engine_name: 'TrendMicro',
                engine_update: '20160928',
                engine_version: '9.740.0.1012',
                method: 'blacklist',
                result: null,
              },
              'TrendMicro-HouseCall': {
                category: 'undetected',
                engine_name: 'TrendMicro-HouseCall',
                engine_update: '20160928',
                engine_version: '9.850.0.1008',
                method: 'blacklist',
                result: null,
              },
              VBA32: {
                category: 'undetected',
                engine_name: 'VBA32',
                engine_update: '20160927',
                engine_version: '3.12.26.4',
                method: 'blacklist',
                result: null,
              },
              VIPRE: {
                category: 'undetected',
                engine_name: 'VIPRE',
                engine_update: '20160928',
                engine_version: '52638',
                method: 'blacklist',
                result: null,
              },
              ViRobot: {
                category: 'undetected',
                engine_name: 'ViRobot',
                engine_update: '20160928',
                engine_version: '2014.3.20.0',
                method: 'blacklist',
                result: null,
              },
              Yandex: {
                category: 'undetected',
                engine_name: 'Yandex',
                engine_update: '20160927',
                engine_version: '5.5.1.3',
                method: 'blacklist',
                result: null,
              },
              Zillya: {
                category: 'undetected',
                engine_name: 'Zillya',
                engine_update: '20160927',
                engine_version: '2.0.0.3070',
                method: 'blacklist',
                result: null,
              },
              Zoner: {
                category: 'undetected',
                engine_name: 'Zoner',
                engine_update: '20160928',
                engine_version: '1.0',
                method: 'blacklist',
                result: null,
              },
              nProtect: {
                category: 'undetected',
                engine_name: 'nProtect',
                engine_update: '20160928',
                engine_version: '2016-09-28.02',
                method: 'blacklist',
                result: null,
              },
            },
            stats: {
              'confirmed-timeout': 0,
              failure: 0,
              harmless: 0,
              malicious: 0,
              suspicious: 0,
              timeout: 0,
              'type-unsupported': 0,
              undetected: 55,
            },
            status: 'completed',
          },
          id:
            'f-2c1be40c41023b35e356267ffd49bab37571a89f6858757b4c5586aa8fd2487a-1475048277',
          links: {
            self:
              'https://www.virustotal.com/api/v3/analyses/f-2c1be40c41023b35e356267ffd49bab37571a89f6858757b4c5586aa8fd2487a-1475048277',
          },
          type: 'analysis',
        },
        meta: {
          file_info: {
            md5: '66424e91c022a925c46565d3ace31ffb',
            sha1: '0b58d9af595450e0ebb020b3f58034b453d56098',
            sha256:
              '2c1be40c41023b35e356267ffd49bab37571a89f6858757b4c5586aa8fd2487a',
            size: 70,
          },
        },
      },
    }
  },
  props: {
    upload_filename: String,
    results: Object,
  },
}
</script>

<style scoped>
.result-subtitle {
  font-size: 1.2em;
}
</style>
