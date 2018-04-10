#include <eosiolib/eosio.hpp>
#include <eosiolib/asset.hpp>
#include <eosiolib/print.hpp>

using namespace eosio;

class hello : public eosio::contract {
  public:
      using contract::contract;

      /// @abi action 
      void hi( account_name user ) {
         print( "Hello, ", name{user} );
      }
      /// @abi action
      void add(account_name user) {
         i++;
         print("i:", i, name{user}); // print 1 every action
      }
      /// @abi action
      void store(account_name user) {
          tabs t(_self, user);
          uint64_t i = 1;
          auto it = t.find(i);
          if (it == t.end()) {
            print("insert");
            t.emplace(_self, [&]( auto& s ) {
                s.j++;
            });
          } else {
            print("after insert");
            //t.modify( it, user, [&]( auto& s ) {
            //    s.j++;
            //});
            //print("after j:", it.j, name{user});
          }
      }
      /// @abi action
      void view(account_name user) {
          tabs t(_self, user);
          uint64_t i = 1;
          auto it = t.find(i);
	  if (it == t.end()) {
            print("not found");
          } else {
            print("found");
          }
      }
  private:
     int i;
     struct tab {
        uint64_t i;
        uint64_t j;
        uint64_t primary_key()const {
            return i;
        }
     };
     typedef eosio::multi_index<N(tabs), tab> tabs;
};

EOSIO_ABI( hello, (hi)(add)(store)(view))
