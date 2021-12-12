import React from 'react';

type Props = {
  content: string;
} & typeof defaultProps;

const defaultProps = {
  content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas sodales leo vitae nulla vehicula consequat. Ut quis nunc consectetur, Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas sodales leo vitae nulla vehicula consequat. Ut quis nunc consectetur,"
};

export const Item = (props: Props): JSX.Element => {
  return (
    <div
      style={{
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        border: 'thin solid white',
        overflow: 'auto',
      }}
    >
      {
        props.content
      }
    </div>
  );
}
Item.defaultProps = defaultProps;
