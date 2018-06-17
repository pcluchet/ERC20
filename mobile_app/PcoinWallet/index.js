import React, { Component } from "react";
import {
    AppRegistry,
    StyleSheet,
    ActivityIndicator, // import des composants
    TouchableOpacity,
    Text,
    TextInput,
    View,
} from 'react-native'


const APIURL = "http://192.168.0.2:8000";

export const getUserBalance = (username) => {
    //return "hi";

    return fetch(`${APIURL}`, {
        method: 'POST',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          Transaction: 'balanceOf',
          Id: username,
          TokenOwner: username,
        }),
        })
        .then((response) => response.json())
        .then((responseJson) => {

            console.log("DEBUG: json :" + responseJson);
            console.log("DEBUG: json :" + responseJson.name);

            //this.setState({ isLoading: false })
            //this.setState({ name: responseJson.name })
            return responseJson;
        })
        .catch((error) => {
            console.error(error);
        });

}

export const TransferTokens = (username, to, amount) => {
    //return "hi";


    return fetch(`${APIURL}`, {
        method: 'POST',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          Transaction: 'transfer',
          Id: username,
          To: to,
          Tokens: amount,
        }),
        })
        .then((response) => response.json())
        .then((responseJson) => {

            console.log("DEBUG: json :" + responseJson);
            console.log("DEBUG: json :" + responseJson.name);

            //this.setState({ isLoading: false })
            //this.setState({ name: responseJson.name })
            return responseJson;
        })
        .catch((error) => {
            console.error(error);
        });

}





class App extends Component {



    constructor(props) {
        super(props)

        this.state = {
            name: '', // nom de la bière
            username: '', // nom de la bière
            balance: '0', // nom de la bière
            transferamount: '0', // nom de la bière
            transferto: '', // nom de la bière
            description: '', // sa description
            BalanceIsLoading: false, // la requête API est-elle en cours ?
            UserListIsLoading: false // la requête API est-elle en cours ?
        }
    }


    ft_balanceOfSafe = (json) => {
        if (json.result == 200)
            return json.body;
        else
            return "0"
    }


    ft_get = () => {
        this.setState({ BalanceisLoading: true })
        this.setState({ name: "blah" })
        getUserBalance(this.state.username).then(json => this.setState({
            balance: this.ft_balanceOfSafe(json),
            BalanceisLoading: false // la requête est terminée
        }))
            .catch(error => console.error(error))

        //this.setState({ isLoading: false })
        //this.setState({ name: k })

    }

    refresh_balance = () => {
        this.ft_get();
    }

    ft_transfer = () => {
        this.setState({ BalanceisLoading: true })
        TransferTokens(this.state.username, this.state.transferto, this.state.transferamount)
        .then(json => {

            console.log("DEBUG: jsontransfer :" + json);
            console.log("DEBUG: jsontransferst :" + json.result);
            this.refresh_balance();
            //balance: this.ft_balanceOfSafe(json),

    }
    )
            .catch(error => console.error(error))

        //this.setState({ isLoading: false })
        //this.setState({ name: k })

    }


    render() {
        var Balance = '';
        if (this.state.BalanceisLoading) {
            Balance = <ActivityIndicator /> // si requête en cours, on affiche un spinner
        }
        else {
            Balance =
                <Text style={styles.balancevalue}>
                    {this.state.balance}
                </Text>
        }

        return (
            <View style={styles.container}>
                <TextInput
                    style={{ height: 40, borderColor: 'gray', borderWidth: 1 }}
                    onChangeText={(username) => this.setState({ username })}
                    value={this.state.username}
                />
                <View style={styles.BalanceContainer}>
                    <Text style={styles.name}>
                        Balance :
                    </Text>
                    <View style={styles.Balancevalue}>
                        {Balance}
                    </View>

                    <TouchableOpacity // on ajoute un "bouton" qui requête une autre bière aléatoire
                        onPress={this.ft_get}
                        style={styles.button}
                    >
                        <Text>Refresh</Text>
                    </TouchableOpacity>
                </View>

                <View style={styles.SendContainer}>
                    <Text style={styles.name}>
                        Transfer tokens :
                    </Text>


                <View style={styles.SideBySide}>
                    <Text style={styles.name}>
                        Amount :
                    </Text>
                    <TextInput
                        style={{ height: 40, borderColor: 'gray', borderWidth: 1 }}
                        onChangeText={(transferamount) => this.setState({ transferamount })}
                        value={this.state.transferamount}
                    />

                </View>

                <View style={styles.SideBySide}>
                    <Text style={styles.name}>
                        To :
                    </Text>
                    <TextInput
                        style={{ height: 40, borderColor: 'gray', borderWidth: 1 }}
                        onChangeText={(transferto) => this.setState({ transferto })}
                        value={this.state.transferto}
                    />


                </View>

                    <TouchableOpacity // on ajoute un "bouton" qui requête une autre bière aléatoire
                        onPress={this.ft_transfer}
                        style={styles.buttonsend}
                    >
                        <Text>Send</Text>
                    </TouchableOpacity>
                </View>


            </View>
        )

    }
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        //justifyContent: 'center',
        //alignItems: 'center',
        flexDirection: 'column',
        backgroundColor: '#F5FCFF',
    },
    // ajout de styles divers
    BalanceContainer: {
        height: 34,
        flex: 1,
        flexDirection: 'row',
        justifyContent: 'space-between',
    },
    SendContainer: {
        flex: 1,
        flexDirection: 'column',
    },


    SideBySide: {
        height: 34,
        flex: 1,
        flexDirection: 'row',
        justifyContent: 'space-between',
    },

    name: {
        fontSize: 18,
        fontWeight: '700',
        marginBottom: 10,
    },
    balancevalue: {

        alignItems: 'center',
        fontSize: 18,
        fontWeight: '700',
    },
    description: {
        marginBottom: 10,
    },
    button: {
        height: 30,
        borderWidth: 1,
        backgroundColor: 'green',
        borderRadius: 3,
        padding: 5,
        justifyContent: 'center',
        alignItems: 'center',
    },
    buttonsend: {
        height: 30,
        borderWidth: 1,
        backgroundColor: 'lightblue',
        borderRadius: 3,
        padding: 5,
        justifyContent: 'center',
        alignItems: 'center',
    }

})

AppRegistry.registerComponent("PcoinWallet", () => App);