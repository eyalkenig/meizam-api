import auth0, { Auth0DecodedHash } from 'auth0-js'
import authConfig from '../../auth_config.json'
import Router from '@/router'

const webAuth = new auth0.WebAuth({
  domain: authConfig.domain,
  redirectUri: `${window.location.origin}/callback`,
  clientID: authConfig.clientId,
  responseType: 'id_token',
  scope: 'openid profile email'
})
const loggedInKey = 'loggedIn'
const idTokenKey = 'idToken'
const idTokenExpiryKey = 'idTokenExpiry'

class AuthService {
  login (customState: string, audience: string): void {
    webAuth.authorize({
      state: customState,
      audience: audience
    })
  }
  isAuthenticated (): boolean {
    return localStorage.getItem(loggedInKey) === 'true' && this.isAccessTokenValid()
  }
  localLogin (authResult: Auth0DecodedHash) {
    localStorage.setItem(loggedInKey, 'true')
    localStorage.setItem(idTokenKey, authResult.idToken || '')
    const idTokenExpiry = (authResult.idTokenPayload.exp || 0) * 1000
    localStorage.setItem(idTokenExpiryKey, idTokenExpiry.toString())
  }
  handleAuthentication (): Promise<string> {
    return new Promise((resolve, reject) => {
      webAuth.parseHash((err, authResult) => {
        if (err) {
          reject(err)
        } else {
          if (authResult != null) {
            this.localLogin(authResult)
            resolve(authResult.idToken)
          } else {
            // already logged in
          }
        }
      })
    })
  }
  isAccessTokenValid (): boolean {
    const idToken = localStorage.getItem(idTokenKey)
    const idTokenExpiry = localStorage.getItem(idTokenExpiryKey)
    if (!idToken || !idTokenExpiry) {
      return false
    }
    return (Date.now() < parseInt(idTokenExpiry, 10))
  }
  async getAccessToken (): Promise<string> {
    return new Promise((resolve) => {
      if (this.isAccessTokenValid()) {
        resolve(localStorage.getItem(idTokenKey) || '')
      } else {
        this.renewTokens().then(authResult => {
          resolve(localStorage.getItem(idTokenKey) || '')
        }).catch(e => {
          Router.push('/login')
        })
      }
    })
  }
  renewTokens (): Promise<Auth0DecodedHash> {
    return new Promise((resolve, reject) => {
      if (localStorage.getItem(loggedInKey) !== 'true') {
        return reject(new Error('Not logged in'))
      }

      webAuth.checkSession({}, (err, authResult) => {
        if (err) {
          reject(err)
        } else {
          this.localLogin(authResult)
          resolve(authResult)
        }
      })
    })
  }
}

export default new AuthService()
