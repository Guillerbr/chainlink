import React, { useState, useEffect } from 'react'
import { connect, MapDispatchToProps, MapStateToProps } from 'react-redux'
import { RouteComponentProps } from '@reach/router'
import {
  createStyles,
  withStyles,
  Theme,
  WithStyles,
} from '@material-ui/core/styles'
import Grid from '@material-ui/core/Grid'
import Typography from '@material-ui/core/Typography'
import Paper from '@material-ui/core/Paper'
import List from '../../../components/Admin/Operators/List'
import { ChangePageEvent } from '../../../components/Table'
import { fetchOperators } from '../../../actions/operators'
import { State as AppState } from '../../../reducers'

const EMPTY_MSG =
  "We couldn't find any results for your search query. Try again with the job id, run id, requester, requester id or transaction hash"

const styles = ({ breakpoints, spacing }: Theme) =>
  createStyles({
    container: {
      overflow: 'hidden',
      padding: spacing.unit * 2,
      [breakpoints.up('sm')]: {
        padding: spacing.unit * 3,
      },
    },
    paper: {
      padding: spacing.unit * 2,
      [breakpoints.up('sm')]: {
        padding: spacing.unit * 3,
      },
    },
  })

interface OwnProps {}

interface StateProps {
  authenticated: boolean
}

interface DispatchProps {
  fetchOperators: () => void
}

interface Props
  extends WithStyles<typeof styles>,
    RouteComponentProps,
    StateProps,
    DispatchProps,
    OwnProps {}

export const Index: React.FC<Props> = ({ classes, fetchOperators }) => {
  const [currentPage, setCurrentPage] = useState(0)
  const onChangePage = (_event: ChangePageEvent, page: number) => {
    setCurrentPage(page)
    fetchOperators(query, page + 1, rowsPerPage)
  }

  useEffect(() => {
    fetchOperators()
  }, [])

  return (
    <Grid
      container
      spacing={24}
      alignItems="center"
      className={classes.container}
    >
      <Grid item xs={12}>
        <Paper className={classes.paper}>
          <Typography variant="h3" gutterBottom>
            Operators
          </Typography>

          <List
            currentPage={currentPage}
            operators={operators}
            count={count}
            onChangePage={onChangePage}
            emptyMsg={EMPTY_MSG}
          />
        </Paper>
      </Grid>
    </Grid>
  )
}

const mapStateToProps: MapStateToProps<
  StateProps,
  OwnProps,
  AppState
> = state => {
  return {
    authenticated: state.adminAuth.allowed,
  }
}

const mapDispatchToProps: MapDispatchToProps<DispatchProps, OwnProps> = {
  fetchOperators,
}

export const ConnectedIndex = connect(
  mapStateToProps,
  mapDispatchToProps,
)(Index)

export default withStyles(styles)(ConnectedIndex)
