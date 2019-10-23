import { Action, Dispatch } from 'redux'
import { ThunkAction } from 'redux-thunk'
import { State as AppState } from '../reducers'
import * as api from '../api'

export function fetchOperators(): ThunkAction<
  Promise<void>,
  AppState,
  void,
  Action<string>
> {
  return (dispatch: Dispatch) => {
    return api
      .getOperators()
      .then(operators => {
        dispatch({ type: 'FETCH_OPERATORS_SUCCEEDED', data: operators })
      })
      .catch(error => {
        switch (error.constructor) {
          case UnauthorizedError:
            dispatch({ type: 'ADMIN_SIGNOUT_SUCCEEDED' })
            break
          default:
            dispatch({ type: 'FETCH_OPERATORS_FAILED' })
            break
        }
      })
  }
}
