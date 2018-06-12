import React, { Component } from "react";
import {
    AppRegistry,
    StyleSheet,
    ActivityIndicator, // import des composants
    TouchableOpacity,
    Text,
    View,
} from 'react-native'


export const getMoviesFromApiAsync = () => {
    //return "hi";

    return fetch('http://192.168.0.11/erc20/index.php')
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

export const getUserList = () => {
    //return "hi";

    return fetch('http://192.168.0.11/erc20/userlist.php')
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

        // la state de notre composant est utilisé pour
        // stocker quelques infos renvoyées par l'API
        this.state = {
            name: '', // nom de la bière
            balance: '0', // nom de la bière
            description: '', // sa description
            BalanceIsLoading: false, // la requête API est-elle en cours ?
            UserListIsLoading: false // la requête API est-elle en cours ?
        }
    }




    ft_get = () => {
        this.setState({ BalanceisLoading: true })
        this.setState({ name: "blah" })
        getMoviesFromApiAsync().then(json => this.setState({
            balance: json.name,
            BalanceisLoading: false // la requête est terminée
        }))
            .catch(error => console.error(error))

        //this.setState({ isLoading: false })
        //this.setState({ name: k })

    }

    render() {
        var Balance = '';
        if (this.state.BalanceisLoading) {
            Balance = <ActivityIndicator /> // si requête en cours, on affiche un spinner
        }
        else 
        {
            Balance = 
                    <Text style={styles.balancevalue}>
                        {this.state.balance}
                    </Text>
        }

        return (
            <View style={styles.container}>
                <View style={styles.BalanceContainer}>
                    <Text style={styles.name}>
                        Balance : 
                    </Text>


                <View style={styles.BalanceContainer}>
                    {Balance}
                    </View>

                    <TouchableOpacity // on ajoute un "bouton" qui requête une autre bière aléatoire
                        onPress={this.ft_get}
                        style={styles.button}
                    >
                        <Text>Actualiser</Text>
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
        backgroundColor: '#F5FCFF',
    },
    // ajout de styles divers
    BalanceContainer: {
        margin: 15,
        height : 34,
        flex: 1,
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'center',
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
    }
})

AppRegistry.registerComponent("PcoinWallet", () => App);