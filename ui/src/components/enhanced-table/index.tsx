import {
  TableHead,
  TableRow,
  TableCell,
  TableSortLabel,
  Box,
} from '@mui/material'
import { visuallyHidden } from '@mui/utils'

export function descendingComparator<T>(a: T, b: T, orderBy: keyof T) {
  if (b[orderBy] < a[orderBy]) {
    return -1
  }
  if (b[orderBy] > a[orderBy]) {
    return 1
  }
  return 0
}

export type Order = 'asc' | 'desc'

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function getComparator<Key extends keyof any>(
  order: Order,
  orderBy: Key
): (
  a: { [key in Key]: number | string | boolean },
  b: { [key in Key]: number | string | boolean }
) => number {
  return order === 'desc'
    ? (a, b) => descendingComparator(a, b, orderBy)
    : (a, b) => -descendingComparator(a, b, orderBy)
}

export interface EnhancedHeadCell<T> {
  disablePadding: boolean
  id: keyof T
  label: string
  numeric: boolean
  align?: 'left' | 'center' | 'right'
}

export interface EnhancedTableHeadProps<T> {
  headCells: readonly EnhancedHeadCell<T>[]
  onRequestSort: (event: React.MouseEvent<unknown>, property: keyof T) => void
  order: Order
  orderBy: string
}

export function EnhancedTableHead<T>(props: EnhancedTableHeadProps<T>) {
  const { headCells, order, orderBy, onRequestSort } = props
  const createSortHandler =
    (property: keyof T) => (event: React.MouseEvent<unknown>) => {
      onRequestSort(event, property)
    }

  return (
    <TableHead>
      <TableRow>
        {headCells.map((headCell) => (
          <TableCell
            key={headCell.id.toString()}
            align={headCell.align || 'center'}
            padding={headCell.disablePadding ? 'none' : 'normal'}
            sortDirection={orderBy === headCell.id ? order : false}
          >
            <TableSortLabel
              active={orderBy === headCell.id}
              direction={orderBy === headCell.id ? order : 'asc'}
              onClick={createSortHandler(headCell.id)}
              sx={{
                '& .MuiTableSortLabel-icon': {
                  width: orderBy === headCell.id ? 'auto' : 0,
                  mx: orderBy === headCell.id ? '4px' : 0,
                },
              }}
            >
              {headCell.label}
              {orderBy === headCell.id ? (
                <Box component="span" sx={visuallyHidden}>
                  {order === 'desc' ? 'sorted descending' : 'sorted ascending'}
                </Box>
              ) : null}
            </TableSortLabel>
          </TableCell>
        ))}
        <TableCell align="right">Controls</TableCell>
      </TableRow>
    </TableHead>
  )
}
