import auth0, { Auth0DecodedHash } from 'auth0-js'
import authConfig from '../../auth_config.json'

const webAuth = new auth0.WebAuth({
  domain: authConfig.domain,
  redirectUri: `${window.location.origin}/callback`,
  clientID: authConfig.clientId,
  responseType: 'token id_token',
  scope: 'openid profile email'
})
const loggedInKey = 'loggedIn'
const accessTokenKey = 'accessToken'
const accessTokenExpiryKey = 'accessTokenExpiry'

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
    localStorage.setItem(accessTokenKey, authResult.accessToken || '')
    const accessTokenExpiry = Date.now() + (authResult.expiresIn || 0) * 1000
    localStorage.setItem(accessTokenExpiryKey, accessTokenExpiry.toString())
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
    const accessToken = localStorage.getItem(accessTokenKey)
    const accessTokenExpiry = localStorage.getItem(accessTokenExpiryKey)
    if (!accessToken || !accessTokenExpiry) {
      return false
    }
    return (Date.now() < parseInt(accessTokenExpiry, 10))
  }
  getAccessToken (): Promise<boolean> {
    return new Promise((resolve, reject) => {
      if (this.isAccessTokenValid()) {
        resolve(true)
      } else {
        this.renewTokens().then(authResult => {
          resolve(true)
        }, reject)
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
