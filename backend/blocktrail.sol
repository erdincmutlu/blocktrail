pragma solidity >=0.4.22 <0.6.0;
pragma experimental ABIEncoderV2;

contract Blocktrail {
    struct Person {
	    string name;  // Name of the person
	    string email;  //
	    string gender;  //
	    string pregnant;  //
	    string child;  // Whether has a child or not
	    string chilgender;  // Gender of the child
	    string childage;  // Age range of the child
	    string country; //
    }

    mapping(uint256 => Person) people;

    function getPerson(uint256 id) public view returns (Person memory, bool) {
        if (bytes(people[id].name).length != 0) {
          return (people[id], true);
        } else {
            Person memory person;
            return (person, false);
        }
    }

    function addPerson(uint256 id, Person memory person) public returns(bool) {
        if (bytes(people[id].name).length != 0)  {
            return false;
        }
        people[id] = person;
        return true;
    }
}
