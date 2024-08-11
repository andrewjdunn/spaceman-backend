## Spaceman: Privacy policy - Work in progress

This is an open source go application that stores car parking pool data in GCP firestore database, it is accessed by an Android app.
The source code for both can be found on GitHub under the GNU AGPL license (3.0 or later).

### Data collected by the app

I hereby state, to the best of my knowledge and belief, that I have not programmed this app to collect any personally identifiable information. All data, a list of car parking spaces associated with an account each of which is associated with a list of googles user accounts is stored securly in firestore encypted using keys only availbale to members of the account.This data can be eased at any time by an account admin. No analytics software is present in the app either.

### Explanation of permissions requested in the app

The list of permissions required by the app can be found in the `AndroidManifest.xml` file:

https://github.com/andrewjdunn/spaceman-app/blob/main/app/src/main/AndroidManifest.xml