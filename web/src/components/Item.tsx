import React from 'react';

type RecordContent = {
  String: string;
  Valid: boolean;
}

export type Record = {
  Id: number;
  Date: string; // TODO: Date ?
  Beacon: string;
  Name: string;
  RecordType: number;
  Record: RecordContent;
};

type Props = {
  content: string;
} & typeof defaultProps;

const defaultProps = {
  content: "N/A"
};

export const Item = (props: Props): JSX.Element => {
  console.log(props.content)
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
      {/* TODO: Handle overflowing */}
      {/* TODO: Update props and display accordingly */}
      {
        props.content
      }
    </div>
  );
}
Item.defaultProps = defaultProps;
