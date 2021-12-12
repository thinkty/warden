import React from 'react';
import { Item } from './Item';

type Props = {
  itemLen: string;
} & typeof defaultProps;

const defaultProps = {
  itemLen: '150px',
};

export const Container = (props: Props): JSX.Element => {

  const items: JSX.Element[] = [];
  for (let i = 0; i < 50; i++) {
    items.push(<Item />);
  }

  return (
    <div
      style={{
        margin: '10px',
        display: 'grid',
        maxHeight: '95vh',
        width: '100vw',
        overflow: 'auto',
        gap: '10px',
        gridAutoRows: props.itemLen,
        gridTemplateColumns: `repeat(auto-fit, ${props.itemLen})`,
        justifyContent: 'center',
      }}
    >
      {
        items
      }
    </div>
  );
}
Container.defaultProps = defaultProps;