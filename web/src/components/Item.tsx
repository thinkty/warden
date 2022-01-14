import React from 'react';

type UserAction = {
  label: string;
  value: string;
};

export type SensorRecord = {
  Id: number;
  Date: Date;
  Beacon: string;
  Name: string;
  Record: { String: string, Valid: boolean };
  UserActions: { String: string, Valid: boolean };
};

type Props = {
  records: SensorRecord[];
} & typeof defaultProps;

const defaultProps = {};

export const Item = ({ records }: Props): JSX.Element => {

  return (
    <div
      onClick={() => {
        console.log(records); // TODO:
        console.log(parseUserActions(records[0].UserActions.String)); // TODO:
      }}
      style={{
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'space-between',
        alignItems: 'center',
        border: 'thin solid white',
        overflow: 'auto',
        cursor: 'pointer',
      }}
    >
      {/* TODO: Handle overflowing */}
      {/* TODO: Update props and display accordingly */}
      {/* Item Header */}
      <div
        style={{
          display: 'flex',
          flexDirection: 'column',
          justifyContent: 'center',
          alignItems: 'center',
        }}
      >
        <div
          style={{
            fontSize: 20,
          }}
        >
          { records[0].Name }
        </div>
        <div
          style={{
            fontSize: 13,
          }}
        >
          { records[0].Beacon }
        </div>
      </div>
      {/* Item Content */}
      <div
        style={{
          fontSize: 44,
          fontWeight: 'bold',
        }}
      >
        { records[0].Record.String }
      </div>
      {/* Item Footer */}
      <div
        style={{

        }}
      >
      </div>
    </div>
  );
}
Item.defaultProps = defaultProps;

/**
 * Parse raw user-actions retrieved from the database. The argument is separated
 * into each user-action by commas and each user-action is a colon separated
 * data structure similar to a dictionary. The left side is the label which will
 * be displayed to the user and the right side is the value that will be sent to
 * the server when triggered. If the user value is provided, it will be sent to
 * the server instead of the default value.
 */
function parseUserActions(rawUserActions: string): UserAction[] {
  const impureUserActions = rawUserActions.split(",").map((rawAction) => {
    const dividerIndex = rawAction.indexOf(":");

    // Error
    if (dividerIndex === -1 || dividerIndex === rawAction.length - 1) {
      return null;
    }

    const action: UserAction = {
      label: rawAction.substring(0, dividerIndex),
      value: rawAction.substring(dividerIndex + 1)
    };

    return action;
  });

  const userActions: UserAction[] = [];

  for (let i = 0; i < impureUserActions.length; i++) {
    const action = impureUserActions[i];
    if (action != null) {
      userActions.push(action);
    }
  }

  // The code above is pretty shitty, but I will let it stay bc it hurts my ego.
  return userActions;
}
