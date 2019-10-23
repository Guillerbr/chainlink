import React from 'react'
import Paper from '@material-ui/core/Paper'
import Hidden from '@material-ui/core/Hidden'
import Table, { ChangePageEvent } from '../../Table'
import { LinkColumn, TextColumn, TimeAgoColumn } from '../../Table/TableCell'

const HEADERS = ['Name', 'URL', 'Created At']

const buildNameCol = (operator: ChainlinkNode): LinkColumn => {
  return {
    type: 'link',
    text: operator.name,
    to: `/job-runs/${operator.id}`,
  }
}

type UrlColumn = LinkColumn | TextColumn

const buildUrlCol = (operator: ChainlinkNode): UrlColumn => {
  if (operator.url) {
    return {
      type: 'link',
      text: operator.url,
      to: operator.url,
    }
  }

  return { type: 'text', text: '-' }
}

const buildCreatedAtCol = (operator: ChainlinkNode): TimeAgoColumn => {
  return {
    type: 'time_ago',
    text: operator.createdAt,
  }
}

const rows = (
  operators: ChainlinkNode[],
): [LinkColumn, UrlColumn, TimeAgoColumn][] => {
  return operators.map(o => {
    return [buildNameCol(o), buildUrlCol(o), buildCreatedAtCol(o)]
  })
}

interface Props {
  currentPage: number
  onChangePage: (event: ChangePageEvent, page: number) => void
  operators?: ChainlinkNode[]
  count?: number
  emptyMsg?: string
  className?: string
}

const List = ({
  operators = [],
  count,
  currentPage,
  className,
  onChangePage,
  emptyMsg,
}: Props) => {
  return (
    <Paper className={className}>
      <Hidden xsDown>
        <Table
          headers={HEADERS}
          currentPage={currentPage}
          rows={rows(operators)}
          count={count}
          onChangePage={onChangePage}
          emptyMsg={emptyMsg}
        />
      </Hidden>
    </Paper>
  )
}

export default List
