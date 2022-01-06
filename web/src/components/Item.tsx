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
  record: Record;
} & typeof defaultProps;

const defaultProps = {};

export const Item = ({ record }: Props): JSX.Element => {

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
        record.Record.String
      }
    </div>
  );
}
Item.defaultProps = defaultProps;
