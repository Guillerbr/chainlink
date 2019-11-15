import React from 'react'
import {
  createStyles,
  Theme,
  withStyles,
  WithStyles,
} from '@material-ui/core/styles'
import { Link as RouterLink } from '@reach/router'
import classNames from 'classnames'

const styles = ({ palette }: Theme) =>
  createStyles({
    link: {
      color: palette.primary.main,
      textDecoration: 'none',
    },
  })

interface Props extends WithStyles<typeof styles> {
  children: React.ReactNode | string
  to: string
  className?: string
}

const PROTOCOL = /^https?:\/\//

const Link = ({ to, children, classes, className }: Props) => {
  if (PROTOCOL.test(to)) {
    return <a href={to}>{children}</a>
  }

  return (
    <RouterLink to={to} className={classNames(classes.link, className)}>
      {children}
    </RouterLink>
  )
}

export default withStyles(styles)(Link)
