import axios from 'axios'
import auth from '@/services/auth-service'

class MeizamApi {
  async fetchTeams (): Promise<any> {
    return this.get('/teams')
  }
  private async get (path: string): Promise<any> {
    const accessToken = await auth.getAccessToken()
    const basePath = process.env.VUE_APP_MEIZAM_API_BASE_URL
    return axios.get(`${basePath}${path}`, {
      headers: {
        Authorization: `Bearer ${accessToken}`
      }
    })
  }
}

export default new MeizamApi()
